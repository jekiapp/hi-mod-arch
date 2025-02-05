package logic

import "github.com/jekiapp/hi-mod/internal/model"

type CalculatePriceItf interface {
	GetProductPrice(productID int64) (float64, error)
	GetPromotion(userID, productID int64) (model.PromotionData, error)
}

func CalculatePrice(cart *model.CartData, itf CalculatePriceItf) error {
	items := cart.Items
	totalPrice := float64(0)
	for _, item := range items {

		currentPrice, err := itf.GetProductPrice(item.ProductID)
		if err != nil {
			return err
		}

		promo, err := itf.GetPromotion(cart.UserID, item.ProductID)
		if err != nil {
			return err
		}
		currentPrice -= promo.Discount
		subtotal := currentPrice * float64(item.Quantity)
		totalPrice += subtotal
	}

	cart.TotalPrice = totalPrice
	return nil
}
