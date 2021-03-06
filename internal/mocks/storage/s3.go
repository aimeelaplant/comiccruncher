// Code generated by MockGen. DO NOT EDIT.
// Source: storage/s3.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	bytes "bytes"
	storage "github.com/comiccruncher/comiccruncher/storage"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	gomock "github.com/golang/mock/gomock"
	io "io"
	http "net/http"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockStorage) Download(key string) (*bytes.Reader, error) {
	ret := m.ctrl.Call(m, "Download", key)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockStorageMockRecorder) Download(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockStorage)(nil).Download), key)
}

// UploadFromRemote mocks base method
func (m *MockStorage) UploadFromRemote(remoteURL, remoteDir string) (storage.UploadedImage, error) {
	ret := m.ctrl.Call(m, "UploadFromRemote", remoteURL, remoteDir)
	ret0, _ := ret[0].(storage.UploadedImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFromRemote indicates an expected call of UploadFromRemote
func (mr *MockStorageMockRecorder) UploadFromRemote(remoteURL, remoteDir interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFromRemote", reflect.TypeOf((*MockStorage)(nil).UploadFromRemote), remoteURL, remoteDir)
}

// UploadBytes mocks base method
func (m *MockStorage) UploadBytes(b *bytes.Buffer, remotePathName string) error {
	ret := m.ctrl.Call(m, "UploadBytes", b, remotePathName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadBytes indicates an expected call of UploadBytes
func (mr *MockStorageMockRecorder) UploadBytes(b, remotePathName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadBytes", reflect.TypeOf((*MockStorage)(nil).UploadBytes), b, remotePathName)
}

// MockS3Client is a mock of S3Client interface
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// GetObject mocks base method
func (m *MockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	ret := m.ctrl.Call(m, "GetObject", input)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockS3ClientMockRecorder) GetObject(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3Client)(nil).GetObject), input)
}

// PutObject mocks base method
func (m *MockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	ret := m.ctrl.Call(m, "PutObject", input)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject
func (mr *MockS3ClientMockRecorder) PutObject(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3Client)(nil).PutObject), input)
}

// MockS3Downloader is a mock of S3Downloader interface
type MockS3Downloader struct {
	ctrl     *gomock.Controller
	recorder *MockS3DownloaderMockRecorder
}

// MockS3DownloaderMockRecorder is the mock recorder for MockS3Downloader
type MockS3DownloaderMockRecorder struct {
	mock *MockS3Downloader
}

// NewMockS3Downloader creates a new mock instance
func NewMockS3Downloader(ctrl *gomock.Controller) *MockS3Downloader {
	mock := &MockS3Downloader{ctrl: ctrl}
	mock.recorder = &MockS3DownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3Downloader) EXPECT() *MockS3DownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockS3Downloader) Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (int64, error) {
	varargs := []interface{}{w, input}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Download", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockS3DownloaderMockRecorder) Download(w, input interface{}, options ...interface{}) *gomock.Call {
	varargs := append([]interface{}{w, input}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockS3Downloader)(nil).Download), varargs...)
}

// MockHTTPClient is a mock of HTTPClient interface
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHTTPClientMockRecorder) Get(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), url)
}
