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

type ConvertCartItemToCheckoutItemItf interface {
	GetProductData(productID int64) (model.ProductData, error)
}

func ConvertCartItemToCheckoutItem(cartItems []model.CartItem, itf ConvertCartItemToCheckoutItemItf) ([]model.CheckoutItem, error) {
	checkoutItems := make([]model.CheckoutItem, 0)
	for _, item := range cartItems {
		product, err := itf.GetProductData(item.ProductID)
		if err != nil {
			return nil, err
		}

		subtotal := product.ProductPrice * float64(item.Quantity)

		checkItem := model.CheckoutItem{
			Quantity: item.Quantity,
			Subtotal: subtotal,
		}
		checkoutItems = append(checkoutItems, checkItem)
	}

	return checkoutItems, nil
}
