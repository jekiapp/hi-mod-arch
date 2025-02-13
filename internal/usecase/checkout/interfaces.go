package checkout

import "github.com/jekiapp/hi-mod-arch/internal/model"

//go:generate mockgen -source=interfaces.go -destination=mock_test.go -package=checkout
type IRenderPage interface {
	GetUserInfo(userID int64) (model.UserData, error)
	GetCartFromDB(userID int64) (model.CartData, error)
	GetProductData(productID int64) (model.ProductData, error)
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
}
