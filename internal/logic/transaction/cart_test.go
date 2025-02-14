package transaction

import (
	"github.com/golang/mock/gomock"
	mock_transaction "github.com/jekiapp/hi-mod-arch/internal/logic/transaction/mock"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertCartItemToCheckoutItem(t *testing.T) {
	ctl := gomock.NewController(t)
	itf := mock_transaction.NewMockIConvertCartItemToCheckoutItem(ctl)

	testcases := map[string]struct {
		mock   func() []model.CartItem
		expect []model.CheckoutItem
		err    error
	}{
		"normal": {
			mock: func() []model.CartItem {
				cartItems := []model.CartItem{
					{
						ProductID: 123,
						Quantity:  3,
					},
				}

				itf.EXPECT().GetProductData(int64(123)).Return(model.ProductData{
					ProductID:    123,
					ProductName:  "Shoe",
					ProductPrice: 9000,
				}, nil)

				return cartItems
			},
			expect: []model.CheckoutItem{
				{
					Product: model.ProductData{
						ProductID:    123,
						ProductName:  "Shoe",
						ProductPrice: 9000,
					},
					Quantity: 3,
					Subtotal: 27000,
				},
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			input := tc.mock()

			output, err := ConvertCartItemToCheckoutItem(input, itf)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expect, output)
		})
	}

}
