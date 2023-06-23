// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GenesisEducationKyiv/main-project-delveper/internal/subscription (interfaces: Subscriber)

// Package subscription is a generated GoMock package.
package subscription

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSubscriber is a mock of Subscriber interface.
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber.
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance.
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// SendEmails mocks base method.
func (m *MockSubscriber) SendEmails() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmails")
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmails indicates an expected call of SendEmails.
func (mr *MockSubscriberMockRecorder) SendEmails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmails", reflect.TypeOf((*MockSubscriber)(nil).SendEmails))
}

// Subscribe mocks base method.
func (m *MockSubscriber) Subscribe(arg0 Email) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscriberMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscriber)(nil).Subscribe), arg0)
}
