// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/kube-scheduler-simulator/export (interfaces: PersistentVolumeClaimService)

// Package mock_export is a generated GoMock package.
package mock_export

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/client-go/applyconfigurations/core/v1"
	reflect "reflect"
)

// MockPersistentVolumeClaimService is a mock of PersistentVolumeClaimService interface
type MockPersistentVolumeClaimService struct {
	ctrl     *gomock.Controller
	recorder *MockPersistentVolumeClaimServiceMockRecorder
}

// MockPersistentVolumeClaimServiceMockRecorder is the mock recorder for MockPersistentVolumeClaimService
type MockPersistentVolumeClaimServiceMockRecorder struct {
	mock *MockPersistentVolumeClaimService
}

// NewMockPersistentVolumeClaimService creates a new mock instance
func NewMockPersistentVolumeClaimService(ctrl *gomock.Controller) *MockPersistentVolumeClaimService {
	mock := &MockPersistentVolumeClaimService{ctrl: ctrl}
	mock.recorder = &MockPersistentVolumeClaimServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersistentVolumeClaimService) EXPECT() *MockPersistentVolumeClaimServiceMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockPersistentVolumeClaimService) Apply(arg0 context.Context, arg1 *v10.PersistentVolumeClaimApplyConfiguration) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply
func (mr *MockPersistentVolumeClaimServiceMockRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockPersistentVolumeClaimService)(nil).Apply), arg0, arg1)
}

// Get mocks base method
func (m *MockPersistentVolumeClaimService) Get(arg0 context.Context, arg1 string) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPersistentVolumeClaimServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersistentVolumeClaimService)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockPersistentVolumeClaimService) List(arg0 context.Context) (*v1.PersistentVolumeClaimList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaimList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPersistentVolumeClaimServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPersistentVolumeClaimService)(nil).List), arg0)
}
