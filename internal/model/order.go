package model

type OrderData struct {
	OrderID     int64
	UserID      int64
	OrderAmount float64
	OrderItems  []OrderItem
}
type OrderItem struct {
	ProductID  int64
	Qty        int
	TotalPrice float64
}
