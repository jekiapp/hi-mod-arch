package checkout

import (
	"github.com/golang/mock/gomock"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockrenderPageInterface(ctrl)
	testCase := map[string]struct {
		mock   func() *model.CheckoutPageRequest
		Output interface{}
		err    error
	}{
		"success": {
			mock: func() *model.CheckoutPageRequest {
				userID := int64(1234)
				coupon := "discount"
				req := &model.CheckoutPageRequest{
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
			Output: model.CheckoutPageResponse{
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
		// more test coverage
	}

	uc := renderPageUsecase{mock}
	for name, tc := range testCase {
		t.Run(name, func(t *testing.T) {
			req := tc.mock()

			resp, err := uc.HandlerFunc(req)
			if err != nil {
				assert.Error(t, tc.err)
				return
			}

			assert.Equal(t, tc.Output, resp)
		})
	}
}
