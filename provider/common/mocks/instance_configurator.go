// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/common (interfaces: InstanceConfigurator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	network "github.com/juju/juju/network"
	reflect "reflect"
)

// MockInstanceConfigurator is a mock of InstanceConfigurator interface
type MockInstanceConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceConfiguratorMockRecorder
}

// MockInstanceConfiguratorMockRecorder is the mock recorder for MockInstanceConfigurator
type MockInstanceConfiguratorMockRecorder struct {
	mock *MockInstanceConfigurator
}

// NewMockInstanceConfigurator creates a new mock instance
func NewMockInstanceConfigurator(ctrl *gomock.Controller) *MockInstanceConfigurator {
	mock := &MockInstanceConfigurator{ctrl: ctrl}
	mock.recorder = &MockInstanceConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstanceConfigurator) EXPECT() *MockInstanceConfiguratorMockRecorder {
	return m.recorder
}

// ChangeIngressRules mocks base method
func (m *MockInstanceConfigurator) ChangeIngressRules(arg0 string, arg1 bool, arg2 []network.IngressRule) error {
	ret := m.ctrl.Call(m, "ChangeIngressRules", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeIngressRules indicates an expected call of ChangeIngressRules
func (mr *MockInstanceConfiguratorMockRecorder) ChangeIngressRules(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeIngressRules", reflect.TypeOf((*MockInstanceConfigurator)(nil).ChangeIngressRules), arg0, arg1, arg2)
}

// ConfigureExternalIpAddress mocks base method
func (m *MockInstanceConfigurator) ConfigureExternalIpAddress(arg0 int) error {
	ret := m.ctrl.Call(m, "ConfigureExternalIpAddress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureExternalIpAddress indicates an expected call of ConfigureExternalIpAddress
func (mr *MockInstanceConfiguratorMockRecorder) ConfigureExternalIpAddress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureExternalIpAddress", reflect.TypeOf((*MockInstanceConfigurator)(nil).ConfigureExternalIpAddress), arg0)
}

// DropAllPorts mocks base method
func (m *MockInstanceConfigurator) DropAllPorts(arg0 []int, arg1 string) error {
	ret := m.ctrl.Call(m, "DropAllPorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropAllPorts indicates an expected call of DropAllPorts
func (mr *MockInstanceConfiguratorMockRecorder) DropAllPorts(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropAllPorts", reflect.TypeOf((*MockInstanceConfigurator)(nil).DropAllPorts), arg0, arg1)
}

// FindIngressRules mocks base method
func (m *MockInstanceConfigurator) FindIngressRules() ([]network.IngressRule, error) {
	ret := m.ctrl.Call(m, "FindIngressRules")
	ret0, _ := ret[0].([]network.IngressRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIngressRules indicates an expected call of FindIngressRules
func (mr *MockInstanceConfiguratorMockRecorder) FindIngressRules() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIngressRules", reflect.TypeOf((*MockInstanceConfigurator)(nil).FindIngressRules))
}