package transaction

import "github.com/jekiapp/hi-mod-arch/internal/model"

//go:generate mockgen -source=cart.go -destination=mock/cart.go
type IGetCartData interface {
	GetCartFromDB(userID int64) (model.CartData, error)
}

func GetCartData(userID int64, itf IGetCartData) (model.CartData, error) {
	// ...validate user id (optional)
	cartData, err := itf.GetCartFromDB(userID)
	// ...validate cart
	if err != nil {
		// ...handle error
	}
	return cartData, nil
}

type IConvertCartItemToCheckoutItem interface {
	GetProductData(productID int64) (model.ProductData, error)
}

func ConvertCartItemToCheckoutItem(cartItems []model.CartItem, itf IConvertCartItemToCheckoutItem) ([]model.CheckoutItem, error) {
	checkoutItems := make([]model.CheckoutItem, 0)
	for _, item := range cartItems {
		product, err := itf.GetProductData(item.ProductID)
		if err != nil {
			return nil, err
		}

		subtotal := product.ProductPrice * float64(item.Quantity)

		checkItem := model.CheckoutItem{
			Product:  product,
			Quantity: item.Quantity,
			Subtotal: subtotal,
		}
		checkoutItems = append(checkoutItems, checkItem)
	}

	return checkoutItems, nil
}
