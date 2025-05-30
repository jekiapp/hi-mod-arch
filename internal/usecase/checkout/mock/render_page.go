// Code generated by MockGen. DO NOT EDIT.
// Source: render_page.go
//
// Generated by this command:
//
//	mockgen -source=render_page.go -destination=mock/render_page.go
//

// Package mock_checkout is a generated GoMock package.
package mock_checkout

import (
	reflect "reflect"

	model "github.com/jekiapp/hi-mod-arch/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockiRenderPageRepo is a mock of iRenderPageRepo interface.
type MockiRenderPageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockiRenderPageRepoMockRecorder
}

// MockiRenderPageRepoMockRecorder is the mock recorder for MockiRenderPageRepo.
type MockiRenderPageRepoMockRecorder struct {
	mock *MockiRenderPageRepo
}

// NewMockiRenderPageRepo creates a new mock instance.
func NewMockiRenderPageRepo(ctrl *gomock.Controller) *MockiRenderPageRepo {
	mock := &MockiRenderPageRepo{ctrl: ctrl}
	mock.recorder = &MockiRenderPageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiRenderPageRepo) EXPECT() *MockiRenderPageRepoMockRecorder {
	return m.recorder
}

// GetCartFromDB mocks base method.
func (m *MockiRenderPageRepo) GetCartFromDB(userID int64) (model.CartData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartFromDB", userID)
	ret0, _ := ret[0].(model.CartData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartFromDB indicates an expected call of GetCartFromDB.
func (mr *MockiRenderPageRepoMockRecorder) GetCartFromDB(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartFromDB", reflect.TypeOf((*MockiRenderPageRepo)(nil).GetCartFromDB), userID)
}

// GetProductData mocks base method.
func (m *MockiRenderPageRepo) GetProductData(productID int64) (model.ProductData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductData", productID)
	ret0, _ := ret[0].(model.ProductData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductData indicates an expected call of GetProductData.
func (mr *MockiRenderPageRepoMockRecorder) GetProductData(productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductData", reflect.TypeOf((*MockiRenderPageRepo)(nil).GetProductData), productID)
}

// GetPromotion mocks base method.
func (m *MockiRenderPageRepo) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromotion", coupon, totalPrice)
	ret0, _ := ret[0].(model.PromotionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromotion indicates an expected call of GetPromotion.
func (mr *MockiRenderPageRepoMockRecorder) GetPromotion(coupon, totalPrice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromotion", reflect.TypeOf((*MockiRenderPageRepo)(nil).GetPromotion), coupon, totalPrice)
}

// GetUserInfo mocks base method.
func (m *MockiRenderPageRepo) GetUserInfo(userID int64) (model.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", userID)
	ret0, _ := ret[0].(model.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockiRenderPageRepoMockRecorder) GetUserInfo(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockiRenderPageRepo)(nil).GetUserInfo), userID)
}
