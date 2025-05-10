package post_payment

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/jekiapp/hi-mod-arch/internal/logic/price"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/jekiapp/hi-mod-arch/internal/repository/promo"
	tx_repo "github.com/jekiapp/hi-mod-arch/internal/repository/transaction"
	"github.com/jekiapp/hi-mod-arch/pkg/db"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
)

type createOrderUsecase struct {
	repo iCreateOrderRepo
}

func NewCreateOrderUsecase(dbcli *sql.DB, prodCli, promoCli *http.Client) *createOrderUsecase {
	return &createOrderUsecase{
		repo: &createOrderRepo{
			dbCli:        dbcli,
			productCli:   prodCli,
			promotionCli: promoCli,
		},
	}
}

var ERR_PYM_MISMATCH = fmt.Errorf("payment amount mismatch")

func (uc *createOrderUsecase) HandleMessage(ctx context.Context, input model.PaymentSuccess) (output handler.NsqHandlerResult, err error) {
	err = uc.createOrder(input)
	if err != nil {
		if err == ERR_PYM_MISMATCH {
			output.Finish = true
		} else {
			output.Requeue = time.Second
		}
	}
	return output, err
}

func (uc *createOrderUsecase) createOrder(paymentData model.PaymentSuccess) error {
	totalPrice, err := price.CalculateTotalPrice(paymentData.CouponUsed, paymentData.Items, uc.repo)
	if err != nil {
		return fmt.Errorf("error calculate price: %s", err.Error())
	}

	// validate to avoid coupon fraud
	if totalPrice != paymentData.PaymentAmount {
		return ERR_PYM_MISMATCH
	}

	orderData := convertPaymentDataToOrder(paymentData)

	tx, _ := uc.repo.Begin()

	orderID, err := uc.repo.InsertOrder(tx, orderData)
	if err != nil {
		return fmt.Errorf("error insert order: %s", err.Error())
	}

	for _, item := range orderData.OrderItems {
		err := uc.repo.InsertOrderItem(tx, orderID, item)
		if err != nil {
			return fmt.Errorf("error insert order item: %s", err.Error())
		}
	}

	err = uc.repo.Commit(tx)
	if err != nil {
		return fmt.Errorf("error commit tx : %s", err.Error())
	}

	return nil
}

// Example of object conversion function
// Notice that this function is a static function. This way, this function should be easily refactored to logic package
// when there are other usecase(s) that need this function
func convertPaymentDataToOrder(pymData model.PaymentSuccess) model.OrderData {
	orderItem := make([]model.OrderItem, 0)
	for _, item := range pymData.Items {
		oitem := model.OrderItem{
			ProductID:  item.Product.ProductID,
			Qty:        item.Quantity,
			TotalPrice: item.Subtotal,
		}
		orderItem = append(orderItem, oitem)
	}

	return model.OrderData{
		UserID:      pymData.UserID,
		OrderAmount: pymData.PaymentAmount,
		OrderItems:  orderItem,
	}
}

//go:generate mockgen -source=create_order.go -destination=mock/create_order.go
type iCreateOrderRepo interface {
	db.ITransaction
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
	InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error)
	InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error
}

type createOrderRepo struct {
	dbCli        *sql.DB
	productCli   *http.Client
	promotionCli *http.Client
}

func (uc *createOrderRepo) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}

func (uc *createOrderRepo) Begin() (*sql.Tx, error) {
	return uc.dbCli.Begin()
}

func (uc *createOrderRepo) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (uc *createOrderRepo) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (uc *createOrderRepo) InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error) {
	return tx_repo.InsertOrder(tx, order)
}

func (uc *createOrderRepo) InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error {
	return tx_repo.InsertOrderItem(tx, orderID, order)
}
