package transaction

import (
	"database/sql"
	"github.com/jekiapp/hi-mod-arch/internal/model"
)

func InsertOrder(tx *sql.Tx, order model.OrderData) (int64, error) {
	query := "INSERT into order(user_id,amount) values ($1,$2) RETURNING order_id"
	res, err := tx.Exec(query, order.UserID, order.OrderAmount)

	if err != nil {
		// handle error
	}

	return res.LastInsertId()
}

func InsertOrderItem(tx *sql.Tx, orderID int64, order model.OrderItem) error {
	query := `INSERT into order_item(order_id,product_id,qty,total_price) 
				values($1,$2,$3,$4)`
	_, err := tx.Exec(query, orderID, order.ProductID, order.Qty, order.TotalPrice)

	return err
}
