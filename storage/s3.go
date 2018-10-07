package storage

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/aimeelaplant/comiccruncher/internal/hashutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	UploadFromRemote(remoteUrl string, remoteDir string) (UploadedImage, error)
}

type S3Storage struct {
	httpClient     *http.Client
	s3             *s3.S3           // The s3 storage.
	bucket         string           // The name of the S3 bucket.
	namingStrategy FileNameStrategy // The naming strategy for uploading a file to S3.
}

// A callable used for naming a file.
type FileNameStrategy func(basename string) string

// The uploaded image from S3 with its pathname and md5 hash of the image data.
type UploadedImage struct {
	Pathname string
	MD5Hash  string
}

// Uploads a file from a remote url. The remote file gets temporarily read in memory.
func (storage *S3Storage) UploadFromRemote(remoteFile string, remoteDir string) (UploadedImage, error) {
	var uploadImage UploadedImage
	u, err := url.Parse(remoteFile)
	if err != nil {
		return uploadImage, errors.New(fmt.Sprintf("cannot parse url: %s", err))
	}
	res, err := storage.httpClient.Get(remoteFile)
	if err != nil {
		return uploadImage, errors.New(fmt.Sprintf("error requesting the remote url: %s", err))
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotModified {
		return uploadImage, errors.New(fmt.Sprintf("got bad status code from remote url %s: %d", remoteFile, res.StatusCode))
	}
	// Check if there is not a leading slash in the remoteDir.
	if !strings.HasSuffix(remoteDir, "/") {
		// Add a leading slash. :)
		remoteDir = remoteDir + "/"
	}
	remotePathName := remoteDir + storage.namingStrategy(filepath.Base(u.Path))
	// copy for later.
	b, err := ioutil.ReadAll(res.Body)
	nopCloser := ioutil.NopCloser(bytes.NewBuffer(b))
	defer nopCloser.Close()
	if err != nil {
		return uploadImage, err
	}
	if err := storage.uploadBytes(b, remotePathName); err != nil {
		return uploadImage, errors.New(fmt.Sprintf("could not upload: %s", err))
	}
	md5Hash, err := hashutil.MD5Hash(nopCloser)
	if err != nil {
		return uploadImage, err
	}
	uploadImage.MD5Hash = md5Hash
	uploadImage.Pathname = remotePathName
	return uploadImage, nil
}

func (storage *S3Storage) uploadBytes(b []byte, remotePathName string) error {
	ctx := context.Background()
	timeout := time.Duration(10 * time.Second) // 10 seconds
	ctx, cancelFn := context.WithTimeout(ctx, timeout)
	defer cancelFn()
	if _, err := storage.s3.PutObject(
		&s3.PutObjectInput{
			Bucket:        aws.String(storage.bucket),
			Body:          bytes.NewReader(b),
			ContentType:   aws.String(http.DetectContentType(b)),
			ContentLength: aws.Int64(int64(len(b))),
			Key:           aws.String(remotePathName),
			CacheControl:  aws.String("max-age=2592000"),
		},
	); err != nil {
		return err
	}
	return nil
}

// Uploads an opened file to s3. The caller is responsible for closing the file.
func (storage *S3Storage) upload(file *os.File, remotePathName string) error {
	ctx := context.Background()
	timeout := time.Duration(10 * time.Second) // 10 seconds
	ctx, cancelFn := context.WithTimeout(ctx, timeout)
	defer cancelFn()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	if _, err := file.Read(buffer); err != nil {
		return errors.New(fmt.Sprintf("could not read file: %s", err))
	}
	if _, err := storage.s3.PutObject(
		&s3.PutObjectInput{
			Bucket:        aws.String(storage.bucket),
			Body:          bytes.NewReader(buffer),
			ContentType:   aws.String(http.DetectContentType(buffer)),
			ContentLength: aws.Int64(size),
			Key:           aws.String(remotePathName),
			CacheControl:  aws.String("max-age=2592000"),
		},
	); err != nil {
		return err
	}
	return nil
}

func NewS3StorageFromEnv() (Storage, error) {
	creds := credentials.Value{
		AccessKeyID:     os.Getenv("CC_AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("CC_AWS_SECRET_ACCESS_KEY"),
	}
	ses, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("CC_AWS_REGION")),
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	})
	if err != nil {
		return nil, err
	}
	s3Storage := S3Storage{
		httpClient:     http.DefaultClient,
		s3:             s3.New(ses),
		bucket:         os.Getenv("CC_AWS_BUCKET"),
		namingStrategy: Crc32TimeNamingStrategy(),
	}
	return &s3Storage, nil
}

// Returns the crc32 encoded string of the unix time in nanoseconds plus the file extension
// of the given basename.
func Crc32TimeNamingStrategy() FileNameStrategy {
	return func(basename string) string {
		// Create a new instance every time to make it concurrent-safe.
		crcHasher := crc32.NewIEEE()
		crcHasher.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
		return hex.EncodeToString(crcHasher.Sum(nil)) + filepath.Ext(basename)
	}
}

func NewS3Storage(httpClient *http.Client, s3 *s3.S3, bucket string, strategy FileNameStrategy) Storage {
	return &S3Storage{
		httpClient:     httpClient,
		s3:             s3,
		bucket:         bucket,
		namingStrategy: strategy,
	}
}
