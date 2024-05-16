// Code generated by MockGen. DO NOT EDIT.
// Source: ./permission_event_producer.go
//
// Generated by this command:
//
//	mockgen -source=./permission_event_producer.go -package=evtmocks -destination=../mocks/permission.mock.go -typed PermissionEventProducer
//
// Package evtmocks is a generated GoMock package.
package evtmocks

import (
	context "context"
	reflect "reflect"

	event "github.com/ecodeclub/webook/internal/marketing/internal/event"
	gomock "go.uber.org/mock/gomock"
)

// MockPermissionEventProducer is a mock of PermissionEventProducer interface.
type MockPermissionEventProducer struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionEventProducerMockRecorder
}

// MockPermissionEventProducerMockRecorder is the mock recorder for MockPermissionEventProducer.
type MockPermissionEventProducerMockRecorder struct {
	mock *MockPermissionEventProducer
}

// NewMockPermissionEventProducer creates a new mock instance.
func NewMockPermissionEventProducer(ctrl *gomock.Controller) *MockPermissionEventProducer {
	mock := &MockPermissionEventProducer{ctrl: ctrl}
	mock.recorder = &MockPermissionEventProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionEventProducer) EXPECT() *MockPermissionEventProducerMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockPermissionEventProducer) Produce(ctx context.Context, evt event.PermissionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", ctx, evt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockPermissionEventProducerMockRecorder) Produce(ctx, evt any) *PermissionEventProducerProduceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockPermissionEventProducer)(nil).Produce), ctx, evt)
	return &PermissionEventProducerProduceCall{Call: call}
}

// PermissionEventProducerProduceCall wrap *gomock.Call
type PermissionEventProducerProduceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PermissionEventProducerProduceCall) Return(arg0 error) *PermissionEventProducerProduceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PermissionEventProducerProduceCall) Do(f func(context.Context, event.PermissionEvent) error) *PermissionEventProducerProduceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PermissionEventProducerProduceCall) DoAndReturn(f func(context.Context, event.PermissionEvent) error) *PermissionEventProducerProduceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}