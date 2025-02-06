package model

type CheckoutPageRequest struct {
	UserID      int64
	PromoCoupon string
}

type CheckoutPageResponse struct {
	User       UserData
	Items      []CheckoutItem
	FinalPrice float64
}

type CheckoutItem struct {
	Product  ProductData
	Quantity int
	Subtotal float64
}
