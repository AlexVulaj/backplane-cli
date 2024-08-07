// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/backplane-cli/pkg/info (interfaces: BuildInfoService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	debug "runtime/debug"

	gomock "github.com/golang/mock/gomock"
)

// MockBuildInfoService is a mock of BuildInfoService interface.
type MockBuildInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockBuildInfoServiceMockRecorder
}

// MockBuildInfoServiceMockRecorder is the mock recorder for MockBuildInfoService.
type MockBuildInfoServiceMockRecorder struct {
	mock *MockBuildInfoService
}

// NewMockBuildInfoService creates a new mock instance.
func NewMockBuildInfoService(ctrl *gomock.Controller) *MockBuildInfoService {
	mock := &MockBuildInfoService{ctrl: ctrl}
	mock.recorder = &MockBuildInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildInfoService) EXPECT() *MockBuildInfoServiceMockRecorder {
	return m.recorder
}

// GetBuildInfo mocks base method.
func (m *MockBuildInfoService) GetBuildInfo() (*debug.BuildInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuildInfo")
	ret0, _ := ret[0].(*debug.BuildInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetBuildInfo indicates an expected call of GetBuildInfo.
func (mr *MockBuildInfoServiceMockRecorder) GetBuildInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuildInfo", reflect.TypeOf((*MockBuildInfoService)(nil).GetBuildInfo))
}
