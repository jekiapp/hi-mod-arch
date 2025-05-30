// Code generated by MockGen. DO NOT EDIT.
// Source: create_order.go
//
// Generated by this command:
//
//	mockgen -source=create_order.go -destination=mock/create_order.go
//

// Package mock_post_payment is a generated GoMock package.
package mock_post_payment

import (
	sql "database/sql"
	reflect "reflect"

	model "github.com/jekiapp/hi-mod-arch/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockiCreateOrderRepo is a mock of iCreateOrderRepo interface.
type MockiCreateOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockiCreateOrderRepoMockRecorder
}

// MockiCreateOrderRepoMockRecorder is the mock recorder for MockiCreateOrderRepo.
type MockiCreateOrderRepoMockRecorder struct {
	mock *MockiCreateOrderRepo
}

// NewMockiCreateOrderRepo creates a new mock instance.
func NewMockiCreateOrderRepo(ctrl *gomock.Controller) *MockiCreateOrderRepo {
	mock := &MockiCreateOrderRepo{ctrl: ctrl}
	mock.recorder = &MockiCreateOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiCreateOrderRepo) EXPECT() *MockiCreateOrderRepoMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockiCreateOrderRepo) Begin() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockiCreateOrderRepoMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockiCreateOrderRepo)(nil).Begin))
}

// Commit mocks base method.
func (m *MockiCreateOrderRepo) Commit(tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockiCreateOrderRepoMockRecorder) Commit(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockiCreateOrderRepo)(nil).Commit), tx)
}

// GetPromotion mocks base method.
func (m *MockiCreateOrderRepo) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromotion", coupon, totalPrice)
	ret0, _ := ret[0].(model.PromotionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromotion indicates an expected call of GetPromotion.
func (mr *MockiCreateOrderRepoMockRecorder) GetPromotion(coupon, totalPrice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromotion", reflect.TypeOf((*MockiCreateOrderRepo)(nil).GetPromotion), coupon, totalPrice)
}

// InsertOrder mocks base method.
func (m *MockiCreateOrderRepo) InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrder", tx, order)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOrder indicates an expected call of InsertOrder.
func (mr *MockiCreateOrderRepoMockRecorder) InsertOrder(tx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrder", reflect.TypeOf((*MockiCreateOrderRepo)(nil).InsertOrder), tx, order)
}

// InsertOrderItem mocks base method.
func (m *MockiCreateOrderRepo) InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrderItem", tx, orderID, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrderItem indicates an expected call of InsertOrderItem.
func (mr *MockiCreateOrderRepoMockRecorder) InsertOrderItem(tx, orderID, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrderItem", reflect.TypeOf((*MockiCreateOrderRepo)(nil).InsertOrderItem), tx, orderID, order)
}

// Rollback mocks base method.
func (m *MockiCreateOrderRepo) Rollback(tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockiCreateOrderRepoMockRecorder) Rollback(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockiCreateOrderRepo)(nil).Rollback), tx)
}
