// Code generated by MockGen. DO NOT EDIT.
// Source: utility/internal/repo (interfaces: PaymentGatewayRepoImply)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	entities "utility/internal/entities"

	gomock "github.com/golang/mock/gomock"
	models "gitlab.com/tuneverse/toolkit/models"
)

// MockPaymentGatewayRepoImply is a mock of PaymentGatewayRepoImply interface.
type MockPaymentGatewayRepoImply struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentGatewayRepoImplyMockRecorder
}

// MockPaymentGatewayRepoImplyMockRecorder is the mock recorder for MockPaymentGatewayRepoImply.
type MockPaymentGatewayRepoImplyMockRecorder struct {
	mock *MockPaymentGatewayRepoImply
}

// NewMockPaymentGatewayRepoImply creates a new mock instance.
func NewMockPaymentGatewayRepoImply(ctrl *gomock.Controller) *MockPaymentGatewayRepoImply {
	mock := &MockPaymentGatewayRepoImply{ctrl: ctrl}
	mock.recorder = &MockPaymentGatewayRepoImplyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentGatewayRepoImply) EXPECT() *MockPaymentGatewayRepoImplyMockRecorder {
	return m.recorder
}

// GetAllPaymentGateway mocks base method.
func (m *MockPaymentGatewayRepoImply) GetAllPaymentGateway(arg0 context.Context, arg1 entities.Params, arg2 entities.Pagination, arg3 entities.Validation, arg4 map[string]models.ErrorResponse) ([]entities.PaymentGateway, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaymentGateway", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]entities.PaymentGateway)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllPaymentGateway indicates an expected call of GetAllPaymentGateway.
func (mr *MockPaymentGatewayRepoImplyMockRecorder) GetAllPaymentGateway(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaymentGateway", reflect.TypeOf((*MockPaymentGatewayRepoImply)(nil).GetAllPaymentGateway), arg0, arg1, arg2, arg3, arg4)
}

// GetPaymentGatewayByID mocks base method.
func (m *MockPaymentGatewayRepoImply) GetPaymentGatewayByID(arg0 context.Context, arg1 string) (entities.PaymentGatewayName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentGatewayByID", arg0, arg1)
	ret0, _ := ret[0].(entities.PaymentGatewayName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentGatewayByID indicates an expected call of GetPaymentGatewayByID.
func (mr *MockPaymentGatewayRepoImplyMockRecorder) GetPaymentGatewayByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentGatewayByID", reflect.TypeOf((*MockPaymentGatewayRepoImply)(nil).GetPaymentGatewayByID), arg0, arg1)
}
