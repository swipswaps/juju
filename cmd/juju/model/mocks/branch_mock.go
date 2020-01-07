// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/model (interfaces: BranchCommandAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBranchCommandAPI is a mock of BranchCommandAPI interface
type MockBranchCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBranchCommandAPIMockRecorder
}

// MockBranchCommandAPIMockRecorder is the mock recorder for MockBranchCommandAPI
type MockBranchCommandAPIMockRecorder struct {
	mock *MockBranchCommandAPI
}

// NewMockBranchCommandAPI creates a new mock instance
func NewMockBranchCommandAPI(ctrl *gomock.Controller) *MockBranchCommandAPI {
	mock := &MockBranchCommandAPI{ctrl: ctrl}
	mock.recorder = &MockBranchCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBranchCommandAPI) EXPECT() *MockBranchCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockBranchCommandAPI) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBranchCommandAPIMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBranchCommandAPI)(nil).Close))
}

// HasActiveBranch mocks base method
func (m *MockBranchCommandAPI) HasActiveBranch(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "HasActiveBranch", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasActiveBranch indicates an expected call of HasActiveBranch
func (mr *MockBranchCommandAPIMockRecorder) HasActiveBranch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasActiveBranch", reflect.TypeOf((*MockBranchCommandAPI)(nil).HasActiveBranch), arg0)
}
