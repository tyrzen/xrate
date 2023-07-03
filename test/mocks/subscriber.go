// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/GenesisEducationKyiv/main-project-delveper/internal/subscription"
	"sync"
)

// Ensure, that SubscriptionServiceMock does implement subscription.SubscriptionService.
// If this is not the case, regenerate this file with moq.
var _ subscription.SubscriptionService = &SubscriptionServiceMock{}

// SubscriptionServiceMock is a mock implementation of subscription.SubscriptionService.
//
//	func TestSomethingThatUsesSubscriptionService(t *testing.T) {
//
//		// make and configure a mocked subscription.SubscriptionService
//		mockedSubscriptionService := &SubscriptionServiceMock{
//			SendEmailsFunc: func() error {
//				panic("mock out the SendEmails method")
//			},
//			SubscribeFunc: func(subscriber subscription.Subscriber) error {
//				panic("mock out the Subscribe method")
//			},
//		}
//
//		// use mockedSubscriptionService in code that requires subscription.SubscriptionService
//		// and then make assertions.
//
//	}
type SubscriptionServiceMock struct {
	// SendEmailsFunc mocks the SendEmails method.
	SendEmailsFunc func() error

	// SubscribeFunc mocks the Subscribe method.
	SubscribeFunc func(subscriber subscription.Subscriber) error

	// calls tracks calls to the methods.
	calls struct {
		// SendEmails holds details about calls to the SendEmails method.
		SendEmails []struct {
		}
		// Subscribe holds details about calls to the Subscribe method.
		Subscribe []struct {
			// Subscriber is the subscriber argument value.
			Subscriber subscription.Subscriber
		}
	}
	lockSendEmails sync.RWMutex
	lockSubscribe  sync.RWMutex
}

// SendEmails calls SendEmailsFunc.
func (mock *SubscriptionServiceMock) SendEmails() error {
	if mock.SendEmailsFunc == nil {
		panic("SubscriptionServiceMock.SendEmailsFunc: method is nil but SubscriptionService.SendEmails was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSendEmails.Lock()
	mock.calls.SendEmails = append(mock.calls.SendEmails, callInfo)
	mock.lockSendEmails.Unlock()
	return mock.SendEmailsFunc()
}

// SendEmailsCalls gets all the calls that were made to SendEmails.
// Check the length with:
//
//	len(mockedSubscriptionService.SendEmailsCalls())
func (mock *SubscriptionServiceMock) SendEmailsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSendEmails.RLock()
	calls = mock.calls.SendEmails
	mock.lockSendEmails.RUnlock()
	return calls
}

// Subscribe calls SubscribeFunc.
func (mock *SubscriptionServiceMock) Subscribe(subscriber subscription.Subscriber) error {
	if mock.SubscribeFunc == nil {
		panic("SubscriptionServiceMock.SubscribeFunc: method is nil but SubscriptionService.Subscribe was just called")
	}
	callInfo := struct {
		Subscriber subscription.Subscriber
	}{
		Subscriber: subscriber,
	}
	mock.lockSubscribe.Lock()
	mock.calls.Subscribe = append(mock.calls.Subscribe, callInfo)
	mock.lockSubscribe.Unlock()
	return mock.SubscribeFunc(subscriber)
}

// SubscribeCalls gets all the calls that were made to Subscribe.
// Check the length with:
//
//	len(mockedSubscriptionService.SubscribeCalls())
func (mock *SubscriptionServiceMock) SubscribeCalls() []struct {
	Subscriber subscription.Subscriber
} {
	var calls []struct {
		Subscriber subscription.Subscriber
	}
	mock.lockSubscribe.RLock()
	calls = mock.calls.Subscribe
	mock.lockSubscribe.RUnlock()
	return calls
}