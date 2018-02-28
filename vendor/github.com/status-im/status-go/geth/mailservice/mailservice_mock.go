// Code generated by MockGen. DO NOT EDIT.
// Source: geth/mailservice/mailservice.go

// Package mailservice is a generated GoMock package.
package mailservice

import (
	node "github.com/ethereum/go-ethereum/node"
	whisperv6 "github.com/ethereum/go-ethereum/whisper/whisperv6"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServiceProvider is a mock of ServiceProvider interface
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderMockRecorder
}

// MockServiceProviderMockRecorder is the mock recorder for MockServiceProvider
type MockServiceProviderMockRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockServiceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProvider) EXPECT() *MockServiceProviderMockRecorder {
	return m.recorder
}

// Node mocks base method
func (m *MockServiceProvider) Node() (*node.Node, error) {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(*node.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Node indicates an expected call of Node
func (mr *MockServiceProviderMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockServiceProvider)(nil).Node))
}

// WhisperService mocks base method
func (m *MockServiceProvider) WhisperService() (*whisperv6.Whisper, error) {
	ret := m.ctrl.Call(m, "WhisperService")
	ret0, _ := ret[0].(*whisperv6.Whisper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhisperService indicates an expected call of WhisperService
func (mr *MockServiceProviderMockRecorder) WhisperService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhisperService", reflect.TypeOf((*MockServiceProvider)(nil).WhisperService))
}
