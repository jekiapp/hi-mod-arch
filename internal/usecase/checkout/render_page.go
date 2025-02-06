package checkout

import (
	"database/sql"
	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/domain"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/logic"
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

	cartData, err := logic.GetCartData(req.UserID, uc)
	if err != nil {
		return nil, err
	}

	user, err := uc.GetUserInfo(req.UserID)
	if err != nil {
		return nil, err
	}

	checkItem, err := logic.ConvertCartItemToCheckoutItem(cartData.Items, uc)
	if err != nil {
		return nil, err
	}

	totalPrice, err := logic.CalculateTotalPrice(req.PromoCoupon, checkItem, uc)
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
	return domain.GetUserInfo(uc.cfg, uc.userCli, userID)
}

func (uc renderPageUsecase) GetCartFromDB(userID int64) (model.CartData, error) {
	return domain.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageUsecase) GetProductData(productID int64) (model.ProductData, error) {
	return domain.GetProductByProductID(uc.cfg, uc.userCli, productID)
}

func (uc renderPageUsecase) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return domain.GetPromotionByCoupon(uc.cfg, uc.promotionCli, coupon, totalPrice)
}
