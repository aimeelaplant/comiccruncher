// Code generated by MockGen. DO NOT EDIT.
// Source: cerebro/characterissue.go

// Package mock_cerebro is a generated GoMock package.
package mock_cerebro

import (
	cerebro "github.com/aimeelaplant/comiccruncher/cerebro"
	comic "github.com/aimeelaplant/comiccruncher/comic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCharacterVendorParser is a mock of CharacterVendorParser interface
type MockCharacterVendorParser struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterVendorParserMockRecorder
}

// MockCharacterVendorParserMockRecorder is the mock recorder for MockCharacterVendorParser
type MockCharacterVendorParserMockRecorder struct {
	mock *MockCharacterVendorParser
}

// NewMockCharacterVendorParser creates a new mock instance
func NewMockCharacterVendorParser(ctrl *gomock.Controller) *MockCharacterVendorParser {
	mock := &MockCharacterVendorParser{ctrl: ctrl}
	mock.recorder = &MockCharacterVendorParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterVendorParser) EXPECT() *MockCharacterVendorParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockCharacterVendorParser) Parse(sources []*comic.CharacterSource) (cerebro.CharacterVendorInfo, error) {
	ret := m.ctrl.Call(m, "Parse", sources)
	ret0, _ := ret[0].(cerebro.CharacterVendorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockCharacterVendorParserMockRecorder) Parse(sources interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockCharacterVendorParser)(nil).Parse), sources)
}
