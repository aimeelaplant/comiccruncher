// Code generated by MockGen. DO NOT EDIT.
// Source: imaging/thumbnail.go

// Package mock_imaging is a generated GoMock package.
package mock_imaging

import (
	bytes "bytes"
	imaging "github.com/comiccruncher/comiccruncher/imaging"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockThumbnailer is a mock of Thumbnailer interface
type MockThumbnailer struct {
	ctrl     *gomock.Controller
	recorder *MockThumbnailerMockRecorder
}

// MockThumbnailerMockRecorder is the mock recorder for MockThumbnailer
type MockThumbnailerMockRecorder struct {
	mock *MockThumbnailer
}

// NewMockThumbnailer creates a new mock instance
func NewMockThumbnailer(ctrl *gomock.Controller) *MockThumbnailer {
	mock := &MockThumbnailer{ctrl: ctrl}
	mock.recorder = &MockThumbnailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThumbnailer) EXPECT() *MockThumbnailerMockRecorder {
	return m.recorder
}

// Resize mocks base method
func (m *MockThumbnailer) Resize(body io.Reader, width, height int) (*bytes.Buffer, error) {
	ret := m.ctrl.Call(m, "Resize", body, width, height)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize
func (mr *MockThumbnailerMockRecorder) Resize(body, width, height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockThumbnailer)(nil).Resize), body, width, height)
}

// MockThumbnailUploader is a mock of ThumbnailUploader interface
type MockThumbnailUploader struct {
	ctrl     *gomock.Controller
	recorder *MockThumbnailUploaderMockRecorder
}

// MockThumbnailUploaderMockRecorder is the mock recorder for MockThumbnailUploader
type MockThumbnailUploaderMockRecorder struct {
	mock *MockThumbnailUploader
}

// NewMockThumbnailUploader creates a new mock instance
func NewMockThumbnailUploader(ctrl *gomock.Controller) *MockThumbnailUploader {
	mock := &MockThumbnailUploader{ctrl: ctrl}
	mock.recorder = &MockThumbnailUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThumbnailUploader) EXPECT() *MockThumbnailUploaderMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockThumbnailUploader) Generate(key string, opts *imaging.ThumbnailOptions) ([]*imaging.ThumbnailResult, error) {
	ret := m.ctrl.Call(m, "Generate", key, opts)
	ret0, _ := ret[0].([]*imaging.ThumbnailResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate
func (mr *MockThumbnailUploaderMockRecorder) Generate(key, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockThumbnailUploader)(nil).Generate), key, opts)
}
