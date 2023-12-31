// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockrepo is a mock of repo interface.
type Mockrepo struct {
	ctrl     *gomock.Controller
	recorder *MockrepoMockRecorder
}

// MockrepoMockRecorder is the mock recorder for Mockrepo.
type MockrepoMockRecorder struct {
	mock *Mockrepo
}

// NewMockrepo creates a new mock instance.
func NewMockrepo(ctrl *gomock.Controller) *Mockrepo {
	mock := &Mockrepo{ctrl: ctrl}
	mock.recorder = &MockrepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepo) EXPECT() *MockrepoMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *Mockrepo) Clear(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear", ctx)
}

// Clear indicates an expected call of Clear.
func (mr *MockrepoMockRecorder) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*Mockrepo)(nil).Clear), ctx)
}

// GetURL mocks base method.
func (m *Mockrepo) GetURL(ctx context.Context, url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", ctx, url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockrepoMockRecorder) GetURL(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*Mockrepo)(nil).GetURL), ctx, url)
}

// SaveHashByURL mocks base method.
func (m *Mockrepo) SaveHashByURL(ctx context.Context, url, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveHashByURL", ctx, url, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveHashByURL indicates an expected call of SaveHashByURL.
func (mr *MockrepoMockRecorder) SaveHashByURL(ctx, url, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveHashByURL", reflect.TypeOf((*Mockrepo)(nil).SaveHashByURL), ctx, url, hash)
}

// Mockgenerator is a mock of generator interface.
type Mockgenerator struct {
	ctrl     *gomock.Controller
	recorder *MockgeneratorMockRecorder
}

// MockgeneratorMockRecorder is the mock recorder for Mockgenerator.
type MockgeneratorMockRecorder struct {
	mock *Mockgenerator
}

// NewMockgenerator creates a new mock instance.
func NewMockgenerator(ctrl *gomock.Controller) *Mockgenerator {
	mock := &Mockgenerator{ctrl: ctrl}
	mock.recorder = &MockgeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockgenerator) EXPECT() *MockgeneratorMockRecorder {
	return m.recorder
}

// GenerateHash mocks base method.
func (m *Mockgenerator) GenerateHash() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHash")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateHash indicates an expected call of GenerateHash.
func (mr *MockgeneratorMockRecorder) GenerateHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHash", reflect.TypeOf((*Mockgenerator)(nil).GenerateHash))
}
