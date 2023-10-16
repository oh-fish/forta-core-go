// Code generated by MockGen. DO NOT EDIT.
// Source: registry/version.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	reflect "reflect"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gomock "github.com/golang/mock/gomock"
)

// MockVersionGetter is a mock of VersionGetter interface.
type MockVersionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockVersionGetterMockRecorder
}

// MockVersionGetterMockRecorder is the mock recorder for MockVersionGetter.
type MockVersionGetterMockRecorder struct {
	mock *MockVersionGetter
}

// NewMockVersionGetter creates a new mock instance.
func NewMockVersionGetter(ctrl *gomock.Controller) *MockVersionGetter {
	mock := &MockVersionGetter{ctrl: ctrl}
	mock.recorder = &MockVersionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionGetter) EXPECT() *MockVersionGetterMockRecorder {
	return m.recorder
}

// Version mocks base method.
func (m *MockVersionGetter) Version(opts *bind.CallOpts) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockVersionGetterMockRecorder) Version(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockVersionGetter)(nil).Version), opts)
}

// MockVersionSetter is a mock of VersionSetter interface.
type MockVersionSetter struct {
	ctrl     *gomock.Controller
	recorder *MockVersionSetterMockRecorder
}

// MockVersionSetterMockRecorder is the mock recorder for MockVersionSetter.
type MockVersionSetterMockRecorder struct {
	mock *MockVersionSetter
}

// NewMockVersionSetter creates a new mock instance.
func NewMockVersionSetter(ctrl *gomock.Controller) *MockVersionSetter {
	mock := &MockVersionSetter{ctrl: ctrl}
	mock.recorder = &MockVersionSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionSetter) EXPECT() *MockVersionSetterMockRecorder {
	return m.recorder
}

// Use mocks base method.
func (m *MockVersionSetter) Use(version string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use", version)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockVersionSetterMockRecorder) Use(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockVersionSetter)(nil).Use), version)
}
