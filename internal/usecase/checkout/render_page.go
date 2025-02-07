package checkout

import (
	"database/sql"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/config"
	cart_domain "github.com/jekiapp/hi-mod/internal/domain/cart"
	product_domain "github.com/jekiapp/hi-mod/internal/domain/product"
	promo_domain "github.com/jekiapp/hi-mod/internal/domain/promo"
	user_domain "github.com/jekiapp/hi-mod/internal/domain/user"
	cart_logic "github.com/jekiapp/hi-mod/internal/logic/cart"
	price_logic "github.com/jekiapp/hi-mod/internal/logic/price"
	"github.com/jekiapp/hi-mod/internal/model"
	"github.com/jekiapp/hi-mod/pkg/handler"
)

type renderPageUsecase struct {
	cfg          *config.Config
	dbCli        *sql.DB
	userCli      *http.Client
	productCli   *http.Client
	promotionCli *http.Client
}

func RenderCheckoutPage(cfg *config.Config, dbCli *sql.DB,
	promotionCli, productCli, userCli *http.Client) handler.GenericHandler {
	return renderPageUsecase{
		cfg:          cfg,
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
	req := input.(*model.CheckoutPageRequest)

	cartData, err := cart_logic.GetCartData(req.UserID, uc)
	if err != nil {
		return nil, err
	}

	user, err := uc.GetUserInfo(req.UserID)
	if err != nil {
		return nil, err
	}

	checkItem, err := cart_logic.ConvertCartItemToCheckoutItem(cartData.Items, uc)
	if err != nil {
		return nil, err
	}

	totalPrice, err := price_logic.CalculateTotalPrice(req.PromoCoupon, checkItem, uc)
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

func (uc renderPageUsecase) GetUserInfo(userID int64) (model.UserData, error) {
	return user_domain.GetUserInfo(uc.cfg, uc.userCli, userID)
}

func (uc renderPageUsecase) GetCartFromDB(userID int64) (model.CartData, error) {
	return cart_domain.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageUsecase) GetProductData(productID int64) (model.ProductData, error) {
	return product_domain.GetProductByProductID(uc.cfg, uc.userCli, productID)
}

func (uc renderPageUsecase) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo_domain.GetPromotionByCoupon(uc.cfg, uc.promotionCli, coupon, totalPrice)
}
