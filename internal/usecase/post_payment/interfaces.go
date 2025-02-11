package post_payment

import (
	"database/sql"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/jekiapp/hi-mod-arch/pkg/db"
)

type createOrderUsecaseItf interface {
	db.ITransaction
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
	InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error)
	InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error
}
