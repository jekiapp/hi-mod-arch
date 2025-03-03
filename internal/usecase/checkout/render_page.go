package checkout

import (
	"database/sql"
	"net/http"

	price_logic "github.com/jekiapp/hi-mod-arch/internal/logic/price"
	tx_logic "github.com/jekiapp/hi-mod-arch/internal/logic/transaction"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	product_repo "github.com/jekiapp/hi-mod-arch/internal/repository/product"
	promo_repo "github.com/jekiapp/hi-mod-arch/internal/repository/promo"
	tx_repo "github.com/jekiapp/hi-mod-arch/internal/repository/transaction"
	user_repo "github.com/jekiapp/hi-mod-arch/internal/repository/user"
)

//go:generate mockgen -source=render_page.go -destination=mock/render_page.go
type renderPageItf interface {
	tx_logic.IGetCartData
	tx_logic.IConvertCartItemToCheckoutItem
	price_logic.ICalculateTotalPrice
	GetUserInfo(userID int64) (model.UserData, error)
}

type renderPageUsecase struct {
	dbCli        *sql.DB
	userCli      *http.Client
	productCli   *http.Client
	promotionCli *http.Client
}

func NewRenderCheckoutPage(dbCli *sql.DB,
	promotionCli, productCli, userCli *http.Client) renderPageUsecase {
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
	return user_repo.GetUserInfo(uc.userCli, userID)
}

func (uc renderPageUsecase) GetCartFromDB(userID int64) (model.CartData, error) {
	return tx_repo.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageUsecase) GetProductData(productID int64) (model.ProductData, error) {
	return product_repo.GetProductByProductID(uc.userCli, productID)
}

func (uc renderPageUsecase) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo_repo.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}
