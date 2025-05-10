package model

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
