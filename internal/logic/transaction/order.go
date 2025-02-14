package transaction

import (
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/nsqio/go-nsq"
)

//go:generate mockgen -source=order.go -destination=mock_test.go -package=transaction
type ICreateOrder interface {
	InsertOrder(data model.OrderData) error
	PublishCreateOrderEvent(msg nsq.Message) error
}

func ConvertPaymentDataToOrder(pymData model.PaymentSuccess) model.OrderData {
	orderItem := make([]model.OrderItem, 0)
	for _, item := range pymData.Items {
		oitem := model.OrderItem{
			ProductID:  item.Product.ProductID,
			Qty:        item.Quantity,
			TotalPrice: item.Subtotal,
		}
		orderItem = append(orderItem, oitem)
	}

	return model.OrderData{
		UserID:      pymData.UserID,
		OrderAmount: pymData.PaymentAmount,
		OrderItems:  orderItem,
	}
}
