// Code generated by MockGen. DO NOT EDIT.
// Source: debug.go

// Package mock_internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDebugLogger is a mock of DebugLogger interface
type MockDebugLogger struct {
	ctrl     *gomock.Controller
	recorder *MockDebugLoggerMockRecorder
}

// MockDebugLoggerMockRecorder is the mock recorder for MockDebugLogger
type MockDebugLoggerMockRecorder struct {
	mock *MockDebugLogger
}

// NewMockDebugLogger creates a new mock instance
func NewMockDebugLogger(ctrl *gomock.Controller) *MockDebugLogger {
	mock := &MockDebugLogger{ctrl: ctrl}
	mock.recorder = &MockDebugLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDebugLogger) EXPECT() *MockDebugLoggerMockRecorder {
	return m.recorder
}

// LogInfo mocks base method
func (m *MockDebugLogger) LogInfo(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogInfo", varargs...)
}

// LogInfo indicates an expected call of LogInfo
func (mr *MockDebugLoggerMockRecorder) LogInfo(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogInfo", reflect.TypeOf((*MockDebugLogger)(nil).LogInfo), v...)
}

// LogDebug mocks base method
func (m *MockDebugLogger) LogDebug(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogDebug", varargs...)
}

// LogDebug indicates an expected call of LogDebug
func (mr *MockDebugLoggerMockRecorder) LogDebug(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDebug", reflect.TypeOf((*MockDebugLogger)(nil).LogDebug), v...)
}

// LogError mocks base method
func (m *MockDebugLogger) LogError(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogError", varargs...)
}

// LogError indicates an expected call of LogError
func (mr *MockDebugLoggerMockRecorder) LogError(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockDebugLogger)(nil).LogError), v...)
}

// LogWarning mocks base method
func (m *MockDebugLogger) LogWarning(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogWarning", varargs...)
}

// LogWarning indicates an expected call of LogWarning
func (mr *MockDebugLoggerMockRecorder) LogWarning(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogWarning", reflect.TypeOf((*MockDebugLogger)(nil).LogWarning), v...)
}

// PrintVersion mocks base method
func (m *MockDebugLogger) PrintVersion(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PrintVersion", varargs...)
}

// PrintVersion indicates an expected call of PrintVersion
func (mr *MockDebugLoggerMockRecorder) PrintVersion(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintVersion", reflect.TypeOf((*MockDebugLogger)(nil).PrintVersion), v...)
}