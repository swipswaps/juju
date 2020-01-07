// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/backups (interfaces: ModelStatusAPI)

// Package backups_test is a generated GoMock package.
package backups_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	base "github.com/juju/juju/api/base"
	names_v3 "gopkg.in/juju/names.v3"
)

// MockModelStatusAPI is a mock of ModelStatusAPI interface
type MockModelStatusAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelStatusAPIMockRecorder
}

// MockModelStatusAPIMockRecorder is the mock recorder for MockModelStatusAPI
type MockModelStatusAPIMockRecorder struct {
	mock *MockModelStatusAPI
}

// NewMockModelStatusAPI creates a new mock instance
func NewMockModelStatusAPI(ctrl *gomock.Controller) *MockModelStatusAPI {
	mock := &MockModelStatusAPI{ctrl: ctrl}
	mock.recorder = &MockModelStatusAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelStatusAPI) EXPECT() *MockModelStatusAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockModelStatusAPI) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockModelStatusAPIMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelStatusAPI)(nil).Close))
}

// ModelStatus mocks base method
func (m *MockModelStatusAPI) ModelStatus(arg0 ...names_v3.ModelTag) ([]base.ModelStatus, error) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModelStatus", varargs...)
	ret0, _ := ret[0].([]base.ModelStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelStatus indicates an expected call of ModelStatus
func (mr *MockModelStatusAPIMockRecorder) ModelStatus(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelStatus", reflect.TypeOf((*MockModelStatusAPI)(nil).ModelStatus), arg0...)
}
