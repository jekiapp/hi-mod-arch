package post_payment

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/jekiapp/hi-mod-arch/internal/model"
	mock_post_payment "github.com/jekiapp/hi-mod-arch/internal/usecase/post_payment/mock"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateOrderUsecase_HandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_post_payment.NewMockiCreateOrderRepo(ctrl)

	tests := []struct {
		name           string
		input          model.PaymentSuccess
		mockSetup      func()
		expectedResult handler.NsqHandlerResult
		expectedErr    error
	}{
		{
			name: "successful order creation",
			input: model.PaymentSuccess{
				UserID:        1,
				PaymentAmount: 100.0,
				CouponUsed:    "TEST10",
				Items: []model.CheckoutItem{
					{
						Product: model.ProductData{
							ProductID: 1,
						},
						Quantity: 1,
						Subtotal: 100.0,
					},
				},
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetPromotion("TEST10", 100.0).
					Return(model.PromotionData{}, nil)

				mockRepo.EXPECT().
					Begin().
					Return(&sql.Tx{}, nil)

				mockRepo.EXPECT().
					InsertOrder(gomock.Any(), gomock.Any()).
					Return(int64(1), nil)

				mockRepo.EXPECT().
					InsertOrderItem(gomock.Any(), int64(1), gomock.Any()).
					Return(nil)

				mockRepo.EXPECT().
					Commit(gomock.Any()).
					Return(nil)
			},
			expectedResult: handler.NsqHandlerResult{},
			expectedErr:    nil,
		},
		{
			name: "payment amount mismatch",
			input: model.PaymentSuccess{
				UserID:        1,
				PaymentAmount: 90.0, // Different from calculated amount
				CouponUsed:    "TEST10",
				Items: []model.CheckoutItem{
					{
						Product: model.ProductData{
							ProductID: 1,
						},
						Quantity: 1,
						Subtotal: 100.0,
					},
				},
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetPromotion("TEST10", 100.0).
					Return(model.PromotionData{}, nil)
			},
			expectedResult: handler.NsqHandlerResult{Finish: true},
			expectedErr:    ERR_PYM_MISMATCH,
		},
		{
			name: "error during order creation",
			input: model.PaymentSuccess{
				UserID:        1,
				PaymentAmount: 100.0,
				CouponUsed:    "TEST10",
				Items: []model.CheckoutItem{
					{
						Product: model.ProductData{
							ProductID: 1,
						},
						Quantity: 1,
						Subtotal: 100.0,
					},
				},
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetPromotion("TEST10", 100.0).
					Return(model.PromotionData{}, errors.New("db error"))
			},
			expectedResult: handler.NsqHandlerResult{Requeue: time.Second},
			expectedErr:    errors.New("error calculate price: db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			uc := &createOrderUsecase{
				repo: mockRepo,
			}

			result, err := uc.HandleMessage(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestConvertPaymentDataToOrder(t *testing.T) {
	input := model.PaymentSuccess{
		UserID:        1,
		PaymentAmount: 100.0,
		Items: []model.CheckoutItem{
			{
				Product: model.ProductData{
					ProductID: 1,
				},
				Quantity: 2,
				Subtotal: 100.0,
			},
		},
	}

	expected := model.OrderData{
		UserID:      1,
		OrderAmount: 100.0,
		OrderItems: []model.OrderItem{
			{
				ProductID:  1,
				Qty:        2,
				TotalPrice: 100.0,
			},
		},
	}

	result := convertPaymentDataToOrder(input)
	assert.Equal(t, expected, result)
}
