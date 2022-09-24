// Code generated by MockGen. DO NOT EDIT.
// Source: ../pinger.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pinger "github.com/echovl/pinger"
	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// GetHost mocks base method.
func (m *MockDB) GetHost(ctx context.Context, hostID int) (*pinger.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", ctx, hostID)
	ret0, _ := ret[0].(*pinger.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHost indicates an expected call of GetHost.
func (mr *MockDBMockRecorder) GetHost(ctx, hostID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockDB)(nil).GetHost), ctx, hostID)
}

// GetHosts mocks base method.
func (m *MockDB) GetHosts(ctx context.Context, limit, skip int) ([]*pinger.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts", ctx, limit, skip)
	ret0, _ := ret[0].([]*pinger.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts.
func (mr *MockDBMockRecorder) GetHosts(ctx, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockDB)(nil).GetHosts), ctx, limit, skip)
}

// RemoveHost mocks base method.
func (m *MockDB) RemoveHost(ctx context.Context, hostID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHost", ctx, hostID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHost indicates an expected call of RemoveHost.
func (mr *MockDBMockRecorder) RemoveHost(ctx, hostID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHost", reflect.TypeOf((*MockDB)(nil).RemoveHost), ctx, hostID)
}

// UpsertHost mocks base method.
func (m *MockDB) UpsertHost(ctx context.Context, host *pinger.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertHost", ctx, host)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertHost indicates an expected call of UpsertHost.
func (mr *MockDBMockRecorder) UpsertHost(ctx, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertHost", reflect.TypeOf((*MockDB)(nil).UpsertHost), ctx, host)
}