package price

import "github.com/jekiapp/hi-mod/internal/model"

type ICalculateTotalPrice interface {
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
}

func CalculateTotalPrice(coupon string, items []model.CheckoutItem, itf ICalculateTotalPrice) (float64, error) {
	totalPrice := float64(0)
	for _, item := range items {
		totalPrice += item.Subtotal
	}

	promo, err := itf.GetPromotion(coupon, totalPrice)
	if err != nil {
		return 0, err
	}

	// ...
	// various validation
	// eligibility logic etc.
	// ...
	if !promo.IsValid {
		return totalPrice, nil
	}

	totalPrice -= promo.Discount

	return totalPrice, nil
}
