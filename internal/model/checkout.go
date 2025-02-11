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

type PaymentSuccess struct {
	UserID        int64
	Items         []CheckoutItem
	CouponUsed    string
	PaymentAmount float64
}
