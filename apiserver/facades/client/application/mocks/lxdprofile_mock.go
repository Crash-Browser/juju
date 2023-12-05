// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/charm/v12 (interfaces: LXDProfiler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charm "github.com/juju/charm/v12"
	gomock "go.uber.org/mock/gomock"
)

// MockLXDProfiler is a mock of LXDProfiler interface.
type MockLXDProfiler struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfilerMockRecorder
}

// MockLXDProfilerMockRecorder is the mock recorder for MockLXDProfiler.
type MockLXDProfilerMockRecorder struct {
	mock *MockLXDProfiler
}

// NewMockLXDProfiler creates a new mock instance.
func NewMockLXDProfiler(ctrl *gomock.Controller) *MockLXDProfiler {
	mock := &MockLXDProfiler{ctrl: ctrl}
	mock.recorder = &MockLXDProfilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfiler) EXPECT() *MockLXDProfilerMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockLXDProfiler) LXDProfile() *charm.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockLXDProfilerMockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockLXDProfiler)(nil).LXDProfile))
}
