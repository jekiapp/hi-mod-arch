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

type renderPageUsecase struct {
	domain IrenderPageDomain
}

type renderPageDomain struct {
	dbCli        *sql.DB
	userCli      *http.Client
	productCli   *http.Client
	promotionCli *http.Client
}

func RenderCheckoutPage(dbCli *sql.DB,
	promotionCli, productCli, userCli *http.Client) handler.GenericHandlerHttp {
	return renderPageUsecase{
		domain: renderPageDomain{
			dbCli:        dbCli,
			productCli:   productCli,
			userCli:      userCli,
			promotionCli: promotionCli,
		},
	}
}

func (uc renderPageUsecase) ObjectAddress() interface{} {
	return &model.CheckoutPageRequest{}
}

func (uc renderPageUsecase) HandlerFunc(input interface{}) (output interface{}, err error) {
	req := input.(*model.CheckoutPageRequest)

	cartData, err := tx_logic.GetCartData(req.UserID, uc.domain)
	if err != nil {
		return nil, err
	}

	user, err := uc.domain.GetUserInfo(req.UserID)
	if err != nil {
		return nil, err
	}

	checkItem, err := tx_logic.ConvertCartItemToCheckoutItem(cartData.Items, uc.domain)
	if err != nil {
		return nil, err
	}

	totalPrice, err := price_logic.CalculateTotalPrice(req.PromoCoupon, checkItem, uc.domain)
	if err != nil {
		return nil, err
	}

	response := model.CheckoutPageResponse{
		User:       user,
		Items:      checkItem,
		FinalPrice: totalPrice,
	}
	return response, nil
}

func (uc renderPageDomain) GetUserInfo(userID int64) (model.UserData, error) {
	return user_domain.GetUserInfo(uc.userCli, userID)
}

func (uc renderPageDomain) GetCartFromDB(userID int64) (model.CartData, error) {
	return tx_domain.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageDomain) GetProductData(productID int64) (model.ProductData, error) {
	return product_domain.GetProductByProductID(uc.userCli, productID)
}

func (uc renderPageDomain) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo_domain.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}
