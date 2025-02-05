package model

type CartData struct {
	UserID     int64
	Items      []CartItem
	TotalPrice float64
}

type CartItem struct {
	ProductID int64
	Quantity  int
}
