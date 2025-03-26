package transaction

import "github.com/jekiapp/hi-mod-arch/internal/model"

//go:generate mockgen -source=cart.go -destination=mock/cart.go
type IConvertCartItemToCheckoutItem interface {
	GetProductData(productID int64) (model.ProductData, error)
}

// Example of object conversion logic
// Currently this function is only used in one usecase
// but this kind of function has a tendency to be reused in other usecase
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
