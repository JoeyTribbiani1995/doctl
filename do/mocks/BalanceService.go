// Code generated by MockGen. DO NOT EDIT.
// Source: balance.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	gomock "github.com/golang/mock/gomock"
)

// MockBalanceService is a mock of BalanceService interface.
type MockBalanceService struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceServiceMockRecorder
}

// MockBalanceServiceMockRecorder is the mock recorder for MockBalanceService.
type MockBalanceServiceMockRecorder struct {
	mock *MockBalanceService
}

// NewMockBalanceService creates a new mock instance.
func NewMockBalanceService(ctrl *gomock.Controller) *MockBalanceService {
	mock := &MockBalanceService{ctrl: ctrl}
	mock.recorder = &MockBalanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBalanceService) EXPECT() *MockBalanceServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBalanceService) Get() (*do.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*do.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBalanceServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBalanceService)(nil).Get))
}
