// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/clsx524/gomock/sample/concurrent (interfaces: Math)

// Package mock_concurrent is a generated GoMock package.
package mock_concurrent

import (
	reflect "reflect"

	gomock "github.com/clsx524/gomock/gomock"
)

// MockMath is a mock of Math interface.
type MockMath struct {
	ctrl     *gomock.Controller
	recorder *MockMathMockRecorder
}

// MockMathMockRecorder is the mock recorder for MockMath.
type MockMathMockRecorder struct {
	mock *MockMath
}

// NewMockMath creates a new mock instance.
func NewMockMath(ctrl *gomock.Controller) *MockMath {
	mock := &MockMath{ctrl: ctrl}
	mock.recorder = &MockMathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMath) EXPECT() *MockMathMockRecorder {
	return m.recorder
}

// Sum mocks base method.
func (m *MockMath) Sum(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sum indicates an expected call of Sum.
func (mr *MockMathMockRecorder) Sum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockMath)(nil).Sum), arg0, arg1)
}
