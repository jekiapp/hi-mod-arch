package checkout

import (
	"context"
	"errors"
	"testing"

	"github.com/jekiapp/hi-mod-arch/internal/model"
	mock_checkout "github.com/jekiapp/hi-mod-arch/internal/usecase/checkout/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHttpGenericHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_checkout.NewMockiRenderPageRepo(ctrl)
	testCase := map[string]struct {
		mock   func() CheckoutPageRequest
		Output CheckoutPageResponse
		err    error
	}{
		"success": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				coupon := "discount"
				req := CheckoutPageRequest{
					UserID:      1234,
					PromoCoupon: coupon,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items: []model.CartItem{
						{
							ProductID: 123,
							Quantity:  3,
						}, {
							ProductID: 789,
							Quantity:  1,
						},
					},
				}, nil)

				mock.EXPECT().GetProductData(int64(123)).Return(model.ProductData{
					ProductID:    123,
					ProductName:  "shoes",
					ProductPrice: 9000,
				}, nil)
				mock.EXPECT().GetProductData(int64(789)).Return(model.ProductData{
					ProductID:    789,
					ProductName:  "shirt",
					ProductPrice: 8000,
				}, nil)

				mock.EXPECT().GetUserInfo(userID).Return(model.UserData{
					UserID:  2323,
					Name:    "Budi",
					Address: "Jakarta",
				}, nil)

				mock.EXPECT().GetPromotion(coupon, gomock.Any()).Return(model.PromotionData{
					IsValid:  true,
					Discount: 5000,
				}, nil)

				return req
			},
			Output: CheckoutPageResponse{
				User: model.UserData{
					UserID:  2323,
					Name:    "Budi",
					Address: "Jakarta",
				},
				Items: []model.CheckoutItem{
					{
						Product: model.ProductData{
							ProductID:    123,
							ProductName:  "shoes",
							ProductPrice: 9000,
						},
						Quantity: 3,
						Subtotal: 27000,
					},
					{
						Product: model.ProductData{
							ProductID:    789,
							ProductName:  "shirt",
							ProductPrice: 8000,
						},
						Quantity: 1,
						Subtotal: 8000,
					},
				},
				FinalPrice: 30000,
			},
		},
		"error_get_cart": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				req := CheckoutPageRequest{
					UserID: userID,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{}, errors.New("db error"))
				return req
			},
			err: errors.New("db error"),
		},
		"error_get_user": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				req := CheckoutPageRequest{
					UserID: userID,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items:  []model.CartItem{},
				}, nil)
				mock.EXPECT().GetUserInfo(userID).Return(model.UserData{}, errors.New("user service error"))
				return req
			},
			err: errors.New("user service error"),
		},
		"empty_cart": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				req := CheckoutPageRequest{
					UserID: userID,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items:  []model.CartItem{},
				}, nil)
				mock.EXPECT().GetUserInfo(userID).Return(model.UserData{
					UserID:  userID,
					Name:    "Budi",
					Address: "Jakarta",
				}, nil)

				return req
			},
			Output: CheckoutPageResponse{
				User: model.UserData{
					UserID:  1234,
					Name:    "Budi",
					Address: "Jakarta",
				},
				Items:      []model.CheckoutItem{},
				FinalPrice: 0,
			},
		},
		"invalid_promo": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				coupon := "invalid"
				req := CheckoutPageRequest{
					UserID:      userID,
					PromoCoupon: coupon,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items: []model.CartItem{
						{
							ProductID: 123,
							Quantity:  1,
						},
					},
				}, nil)

				mock.EXPECT().GetProductData(int64(123)).Return(model.ProductData{
					ProductID:    123,
					ProductName:  "shoes",
					ProductPrice: 9000,
				}, nil)

				mock.EXPECT().GetUserInfo(userID).Return(model.UserData{
					UserID:  userID,
					Name:    "Budi",
					Address: "Jakarta",
				}, nil)

				mock.EXPECT().GetPromotion(coupon, gomock.Any()).Return(model.PromotionData{
					IsValid:  false,
					Discount: 0,
				}, nil)

				return req
			},
			Output: CheckoutPageResponse{
				User: model.UserData{
					UserID:  1234,
					Name:    "Budi",
					Address: "Jakarta",
				},
				Items: []model.CheckoutItem{
					{
						Product: model.ProductData{
							ProductID:    123,
							ProductName:  "shoes",
							ProductPrice: 9000,
						},
						Quantity: 1,
						Subtotal: 9000,
					},
				},
				FinalPrice: 9000,
			},
		},
		"error_get_product": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				req := CheckoutPageRequest{
					UserID: userID,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items: []model.CartItem{
						{
							ProductID: 123,
							Quantity:  1,
						},
					},
				}, nil)

				mock.EXPECT().GetProductData(int64(123)).Return(model.ProductData{}, errors.New("product service error"))
				return req
			},
			err: errors.New("product service error"),
		},
		"error_get_promo": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				coupon := "discount"
				req := CheckoutPageRequest{
					UserID:      userID,
					PromoCoupon: coupon,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items: []model.CartItem{
						{
							ProductID: 123,
							Quantity:  1,
						},
					},
				}, nil)

				mock.EXPECT().GetProductData(int64(123)).Return(model.ProductData{
					ProductID:    123,
					ProductName:  "shoes",
					ProductPrice: 9000,
				}, nil)

				mock.EXPECT().GetUserInfo(userID).Return(model.UserData{
					UserID:  userID,
					Name:    "Budi",
					Address: "Jakarta",
				}, nil)

				mock.EXPECT().GetPromotion(coupon, gomock.Any()).Return(model.PromotionData{}, errors.New("promo service error"))
				return req
			},
			err: errors.New("promo service error"),
		},
		"error_get_product_second_item": {
			mock: func() CheckoutPageRequest {
				userID := int64(1234)
				req := CheckoutPageRequest{
					UserID: userID,
				}

				mock.EXPECT().GetCartFromDB(userID).Return(model.CartData{
					UserID: userID,
					Items: []model.CartItem{
						{
							ProductID: 123,
							Quantity:  1,
						},
						{
							ProductID: 456,
							Quantity:  1,
						},
					},
				}, nil)

				mock.EXPECT().GetProductData(int64(123)).Return(model.ProductData{
					ProductID:    123,
					ProductName:  "shoes",
					ProductPrice: 9000,
				}, nil)
				mock.EXPECT().GetProductData(int64(456)).Return(model.ProductData{}, errors.New("product service error"))
				return req
			},
			err: errors.New("product service error"),
		},
	}

	for name, tc := range testCase {
		t.Run(name, func(t *testing.T) {
			req := tc.mock()
			uc := renderPageUsecase{repo: mock}

			resp, err := uc.HttpGenericHandler(context.Background(), req)
			if err != nil {
				assert.Error(t, tc.err)
				assert.Equal(t, tc.err.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.Output, resp)
		})
	}
}
