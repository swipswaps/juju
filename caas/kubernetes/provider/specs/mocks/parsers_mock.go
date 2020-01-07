// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas/kubernetes/provider/specs (interfaces: PodSpecConverter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	specs "github.com/juju/juju/caas/specs"
)

// MockPodSpecConverter is a mock of PodSpecConverter interface
type MockPodSpecConverter struct {
	ctrl     *gomock.Controller
	recorder *MockPodSpecConverterMockRecorder
}

// MockPodSpecConverterMockRecorder is the mock recorder for MockPodSpecConverter
type MockPodSpecConverterMockRecorder struct {
	mock *MockPodSpecConverter
}

// NewMockPodSpecConverter creates a new mock instance
func NewMockPodSpecConverter(ctrl *gomock.Controller) *MockPodSpecConverter {
	mock := &MockPodSpecConverter{ctrl: ctrl}
	mock.recorder = &MockPodSpecConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodSpecConverter) EXPECT() *MockPodSpecConverterMockRecorder {
	return m.recorder
}

// ToLatest mocks base method
func (m *MockPodSpecConverter) ToLatest() *specs.PodSpecV2 {
	ret := m.ctrl.Call(m, "ToLatest")
	ret0, _ := ret[0].(*specs.PodSpecV2)
	return ret0
}

// ToLatest indicates an expected call of ToLatest
func (mr *MockPodSpecConverterMockRecorder) ToLatest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToLatest", reflect.TypeOf((*MockPodSpecConverter)(nil).ToLatest))
}

// Validate mocks base method
func (m *MockPodSpecConverter) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockPodSpecConverterMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockPodSpecConverter)(nil).Validate))
}
