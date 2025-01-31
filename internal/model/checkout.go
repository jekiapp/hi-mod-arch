package model

type CheckoutPageResponse struct {
	Items []CheckoutItem

	ShippingPrice float64
	FinalPrice    float64
}

type CheckoutItem struct {
	Product  ProductData
	Quantity int
	Subtotal float64
	Discount float64
}
