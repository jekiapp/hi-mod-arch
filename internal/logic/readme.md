# Logic

Logic is for reusable conversion/calculation functions.

## Guidelines

- Write static functions.
- Pass inputs as arguments.
- If logic needs external data, pass an interface.

## Example

```go
type ICalculateTotalPrice interface {
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
}

func CalculateTotalPrice(coupon string, items []model.CheckoutItem, itf ICalculateTotalPrice) (float64, error) {
	total := 0.0
	for _, item := range items {
		total += item.Subtotal
	}
	promo, err := itf.GetPromotion(coupon, total)
	if err != nil {
		return 0, err
	}
	return total - promo.Discount, nil
}
```