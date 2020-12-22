// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nexmoinc/gosrvlib/pkg/testutil (interfaces: TestHTTPResponseWriter)

// Package httputil is a generated GoMock package.
package httputil

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTestHTTPResponseWriter is a mock of TestHTTPResponseWriter interface.
type MockTestHTTPResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockTestHTTPResponseWriterMockRecorder
}

// MockTestHTTPResponseWriterMockRecorder is the mock recorder for MockTestHTTPResponseWriter.
type MockTestHTTPResponseWriterMockRecorder struct {
	mock *MockTestHTTPResponseWriter
}

// NewMockTestHTTPResponseWriter creates a new mock instance.
func NewMockTestHTTPResponseWriter(ctrl *gomock.Controller) *MockTestHTTPResponseWriter {
	mock := &MockTestHTTPResponseWriter{ctrl: ctrl}
	mock.recorder = &MockTestHTTPResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestHTTPResponseWriter) EXPECT() *MockTestHTTPResponseWriterMockRecorder {
	return m.recorder
}

// Header mocks base method.
func (m *MockTestHTTPResponseWriter) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockTestHTTPResponseWriterMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTestHTTPResponseWriter)(nil).Header))
}

// Write mocks base method.
func (m *MockTestHTTPResponseWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockTestHTTPResponseWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTestHTTPResponseWriter)(nil).Write), arg0)
}

// WriteHeader mocks base method.
func (m *MockTestHTTPResponseWriter) WriteHeader(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", arg0)
}

// WriteHeader indicates an expected call of WriteHeader.
func (mr *MockTestHTTPResponseWriterMockRecorder) WriteHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockTestHTTPResponseWriter)(nil).WriteHeader), arg0)
}
