// Code generated by MockGen. DO NOT EDIT.
// Source: app/utils/hash/hash.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHash is a mock of Hash interface.
type MockHash struct {
	ctrl     *gomock.Controller
	recorder *MockHashMockRecorder
}

// MockHashMockRecorder is the mock recorder for MockHash.
type MockHashMockRecorder struct {
	mock *MockHash
}

// NewMockHash creates a new mock instance.
func NewMockHash(ctrl *gomock.Controller) *MockHash {
	mock := &MockHash{ctrl: ctrl}
	mock.recorder = &MockHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHash) EXPECT() *MockHashMockRecorder {
	return m.recorder
}

// ComparePassword mocks base method.
func (m *MockHash) ComparePassword(hashPassword, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", hashPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword.
func (mr *MockHashMockRecorder) ComparePassword(hashPassword, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockHash)(nil).ComparePassword), hashPassword, password)
}

// HashPassword mocks base method.
func (m *MockHash) HashPassword(password string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockHashMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockHash)(nil).HashPassword), password)
}
