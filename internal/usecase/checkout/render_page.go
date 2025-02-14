package checkout

import (
	"database/sql"
	"net/http"

	product_domain "github.com/jekiapp/hi-mod-arch/internal/domain/product"
	promo_domain "github.com/jekiapp/hi-mod-arch/internal/domain/promo"
	tx_domain "github.com/jekiapp/hi-mod-arch/internal/domain/transaction"
	user_domain "github.com/jekiapp/hi-mod-arch/internal/domain/user"
	price_logic "github.com/jekiapp/hi-mod-arch/internal/logic/price"
	tx_logic "github.com/jekiapp/hi-mod-arch/internal/logic/transaction"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
)

//go:generate mockgen -source=render_page.go -destination=mock/render_page.go
type renderPageItf interface {
	GetUserInfo(userID int64) (model.UserData, error)
	GetCartFromDB(userID int64) (model.CartData, error)
	GetProductData(productID int64) (model.ProductData, error)
	GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error)
}

type renderPageUsecase struct {
	dbCli        *sql.DB
	userCli      *http.Client
	productCli   *http.Client
	promotionCli *http.Client
}

func RenderCheckoutPage(dbCli *sql.DB,
	promotionCli, productCli, userCli *http.Client) handler.GenericHandlerHttp {
	return renderPageUsecase{
		dbCli:        dbCli,
		productCli:   productCli,
		userCli:      userCli,
		promotionCli: promotionCli,
	}
}

func (uc renderPageUsecase) ObjectAddress() interface{} {
	return &model.CheckoutPageRequest{}
}

func (uc renderPageUsecase) HandlerFunc(input interface{}) (output interface{}, err error) {
	in := input.(*model.CheckoutPageRequest)
	return renderPage(uc, *in)
}

func renderPage(uc renderPageItf, input model.CheckoutPageRequest) (response model.CheckoutPageResponse, err error) {
	cartData, err := tx_logic.GetCartData(input.UserID, uc)
	if err != nil {
		return response, err
	}

	user, err := uc.GetUserInfo(input.UserID)
	if err != nil {
		return response, err
	}

	checkItem, err := tx_logic.ConvertCartItemToCheckoutItem(cartData.Items, uc)
	if err != nil {
		return response, err
	}

	totalPrice, err := price_logic.CalculateTotalPrice(input.PromoCoupon, checkItem, uc)
	if err != nil {
		return response, err
	}

	response = model.CheckoutPageResponse{
		User:       user,
		Items:      checkItem,
		FinalPrice: totalPrice,
	}
	return response, nil
}

func (uc renderPageUsecase) GetUserInfo(userID int64) (model.UserData, error) {
	return user_domain.GetUserInfo(uc.userCli, userID)
}

func (uc renderPageUsecase) GetCartFromDB(userID int64) (model.CartData, error) {
	return tx_domain.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageUsecase) GetProductData(productID int64) (model.ProductData, error) {
	return product_domain.GetProductByProductID(uc.userCli, productID)
}

func (uc renderPageUsecase) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo_domain.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}
