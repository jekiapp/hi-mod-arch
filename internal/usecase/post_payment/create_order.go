package post_payment

import (
	"database/sql"
	"fmt"
	"github.com/jekiapp/hi-mod-arch/internal/domain/promo"
	tx_domain "github.com/jekiapp/hi-mod-arch/internal/domain/transaction"
	"github.com/jekiapp/hi-mod-arch/internal/logic/price"
	tx_logic "github.com/jekiapp/hi-mod-arch/internal/logic/transaction"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
	"net/http"
	"time"
)

type createOrderUsecase struct {
	dbCli        *sql.DB
	productCli   *http.Client
	promotionCli *http.Client
}

func (uc *createOrderUsecase) HandlerFunc(input interface{}) (output handler.NsqHandlerResult, err error) {
	paymentData := input.(*model.PaymentSuccess)

	totalPrice, err := price.CalculateTotalPrice(paymentData.CouponUsed, paymentData.Items, uc)
	if err != nil {
		output.Requeue = time.Second
		return output, fmt.Errorf("error calculate price: %s", err.Error())
	}

	// validate to avoid coupon fraud
	if totalPrice != paymentData.PaymentAmount {
		output.Finish = true
		return output, fmt.Errorf("payment amount mismatch")
	}

	orderData := tx_logic.ConvertPaymentDataToOrder(*paymentData)

	tx, _ := uc.Begin()

	orderID, err := uc.InsertOrder(tx, orderData)
	if err != nil {
		output.Requeue = time.Second
		return output, fmt.Errorf("error insert order: %s", err.Error())
	}

	for _, item := range orderData.OrderItems {
		err := uc.InsertOrderItem(tx, orderID, item)
		if err != nil {
			output.Requeue = time.Second
			return output, fmt.Errorf("error insert order item: %s", err.Error())
		}
	}

	err = uc.Commit(tx)
	if err != nil {
		output.Requeue = time.Second
		return output, fmt.Errorf("error commit tx : %s", err.Error())
	}

	return output, nil
}

func (uc *createOrderUsecase) ObjectAddress() interface{} {
	return &model.PaymentSuccess{}
}

func (uc *createOrderUsecase) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}

func (uc *createOrderUsecase) Begin() (*sql.Tx, error) {
	return uc.dbCli.Begin()
}

func (uc *createOrderUsecase) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (uc *createOrderUsecase) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (uc *createOrderUsecase) InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error) {
	return tx_domain.InsertOrder(tx, order)
}

func (uc *createOrderUsecase) InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error {
	return tx_domain.InsertOrderItem(tx, orderID, order)
}
