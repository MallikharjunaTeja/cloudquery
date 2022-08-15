// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/k8s/client (interfaces: ServicesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockServicesClient is a mock of ServicesClient interface.
type MockServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicesClientMockRecorder
}

// MockServicesClientMockRecorder is the mock recorder for MockServicesClient.
type MockServicesClientMockRecorder struct {
	mock *MockServicesClient
}

// NewMockServicesClient creates a new mock instance.
func NewMockServicesClient(ctrl *gomock.Controller) *MockServicesClient {
	mock := &MockServicesClient{ctrl: ctrl}
	mock.recorder = &MockServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicesClient) EXPECT() *MockServicesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockServicesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ServiceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServicesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicesClient)(nil).List), arg0, arg1)
}