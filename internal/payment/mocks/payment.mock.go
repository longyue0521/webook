// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -package=paymentmocks -destination=../../mocks/payment.mock.go -typed Service
//
// Package paymentmocks is a generated GoMock package.
package paymentmocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/ecodeclub/webook/internal/payment/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockService) CreatePayment(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", ctx, payment)
	ret0, _ := ret[0].(domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockServiceMockRecorder) CreatePayment(ctx, payment any) *ServiceCreatePaymentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockService)(nil).CreatePayment), ctx, payment)
	return &ServiceCreatePaymentCall{Call: call}
}

// ServiceCreatePaymentCall wrap *gomock.Call
type ServiceCreatePaymentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServiceCreatePaymentCall) Return(arg0 domain.Payment, arg1 error) *ServiceCreatePaymentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServiceCreatePaymentCall) Do(f func(context.Context, domain.Payment) (domain.Payment, error)) *ServiceCreatePaymentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServiceCreatePaymentCall) DoAndReturn(f func(context.Context, domain.Payment) (domain.Payment, error)) *ServiceCreatePaymentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindPaymentByID mocks base method.
func (m *MockService) FindPaymentByID(ctx context.Context, paymentID int64) (domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPaymentByID", ctx, paymentID)
	ret0, _ := ret[0].(domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPaymentByID indicates an expected call of FindPaymentByID.
func (mr *MockServiceMockRecorder) FindPaymentByID(ctx, paymentID any) *ServiceFindPaymentByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPaymentByID", reflect.TypeOf((*MockService)(nil).FindPaymentByID), ctx, paymentID)
	return &ServiceFindPaymentByIDCall{Call: call}
}

// ServiceFindPaymentByIDCall wrap *gomock.Call
type ServiceFindPaymentByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServiceFindPaymentByIDCall) Return(arg0 domain.Payment, arg1 error) *ServiceFindPaymentByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServiceFindPaymentByIDCall) Do(f func(context.Context, int64) (domain.Payment, error)) *ServiceFindPaymentByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServiceFindPaymentByIDCall) DoAndReturn(f func(context.Context, int64) (domain.Payment, error)) *ServiceFindPaymentByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetPaymentChannels mocks base method.
func (m *MockService) GetPaymentChannels(ctx context.Context) []domain.PaymentChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentChannels", ctx)
	ret0, _ := ret[0].([]domain.PaymentChannel)
	return ret0
}

// GetPaymentChannels indicates an expected call of GetPaymentChannels.
func (mr *MockServiceMockRecorder) GetPaymentChannels(ctx any) *ServiceGetPaymentChannelsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentChannels", reflect.TypeOf((*MockService)(nil).GetPaymentChannels), ctx)
	return &ServiceGetPaymentChannelsCall{Call: call}
}

// ServiceGetPaymentChannelsCall wrap *gomock.Call
type ServiceGetPaymentChannelsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServiceGetPaymentChannelsCall) Return(arg0 []domain.PaymentChannel) *ServiceGetPaymentChannelsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServiceGetPaymentChannelsCall) Do(f func(context.Context) []domain.PaymentChannel) *ServiceGetPaymentChannelsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServiceGetPaymentChannelsCall) DoAndReturn(f func(context.Context) []domain.PaymentChannel) *ServiceGetPaymentChannelsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PayByOrderID mocks base method.
func (m *MockService) PayByOrderID(ctx context.Context, oid int64) (domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PayByOrderID", ctx, oid)
	ret0, _ := ret[0].(domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PayByOrderID indicates an expected call of PayByOrderID.
func (mr *MockServiceMockRecorder) PayByOrderID(ctx, oid any) *ServicePayByOrderIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PayByOrderID", reflect.TypeOf((*MockService)(nil).PayByOrderID), ctx, oid)
	return &ServicePayByOrderIDCall{Call: call}
}

// ServicePayByOrderIDCall wrap *gomock.Call
type ServicePayByOrderIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServicePayByOrderIDCall) Return(arg0 domain.Payment, arg1 error) *ServicePayByOrderIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServicePayByOrderIDCall) Do(f func(context.Context, int64) (domain.Payment, error)) *ServicePayByOrderIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServicePayByOrderIDCall) DoAndReturn(f func(context.Context, int64) (domain.Payment, error)) *ServicePayByOrderIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
