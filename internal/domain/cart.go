package domain

import (
	"database/sql"

	"github.com/jekiapp/hi-mod/internal/model"
)

func SelectCartByUserID(db *sql.DB, userID int64) (model.CartData, error) {
	data := model.CartData{}
	query := "SELECT * from cart WHERE user_id=$1"
	rows, err := db.Query(query, userID)
	if err != nil {
		return data, err
	}
	// scan data per row
	// just an example
	rows.Scan(&data)
	// ...
	return data, nil
}
