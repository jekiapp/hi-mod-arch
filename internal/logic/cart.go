package logic

import "github.com/jekiapp/hi-mod/internal/model"

type GetCartDataItf interface {
	GetCartFromDB(userID int64) (model.CartData, error)
}

func GetCartData(userID int64, itf GetCartDataItf) (model.CartData, error) {
	// validate user id
	cartData, err := itf.GetCartFromDB(userID)
	//validate cart
	if err != nil {
		// if err sql no rows
	}
	return cartData, nil
}
