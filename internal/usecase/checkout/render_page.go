package checkout

import (
	"context"
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

type renderPageUsecase struct {
	repo iRenderPageRepo
}

func NewRenderCheckoutPage(dbCli *sql.DB,
	promotionCli, productCli, userCli *http.Client) renderPageUsecase {
	return renderPageUsecase{
		repo: renderPageRepo{
			dbCli:        dbCli,
			productCli:   productCli,
			userCli:      userCli,
			promotionCli: promotionCli,
		},
	}
}

type CheckoutPageRequest struct {
	UserID      int64
	PromoCoupon string
}

type CheckoutPageResponse struct {
	User       model.UserData
	Items      []model.CheckoutItem
	FinalPrice float64
}

func (uc renderPageUsecase) HttpGenericHandler(ctx context.Context, input CheckoutPageRequest) (response CheckoutPageResponse, err error) {
	cartData, err := uc.repo.GetCartFromDB(input.UserID)
	if err != nil {
		return response, err
	}

	user, err := uc.repo.GetUserInfo(input.UserID)
	if err != nil {
		return response, err
	}

	checkItem, err := tx_logic.ConvertCartItemToCheckoutItem(cartData.Items, uc.repo)
	if err != nil {
		return response, err
	}

	totalPrice, err := price_logic.CalculateTotalPrice(input.PromoCoupon, checkItem, uc.repo)
	if err != nil {
		return response, err
	}

	response = CheckoutPageResponse{
		User:       user,
		Items:      checkItem,
		FinalPrice: totalPrice,
	}
	return response, nil
}

//go:generate mockgen -source=render_page.go -destination=mock/render_page.go
type iRenderPageRepo interface {
	tx_logic.IConvertCartItemToCheckoutItem
	price_logic.ICalculateTotalPrice

	GetCartFromDB(userID int64) (model.CartData, error)
	GetUserInfo(userID int64) (model.UserData, error)
}

type renderPageRepo struct {
	dbCli        *sql.DB
	userCli      *http.Client
	productCli   *http.Client
	promotionCli *http.Client
}

func (uc renderPageRepo) GetUserInfo(userID int64) (model.UserData, error) {
	return user_repo.GetUserInfo(uc.userCli, userID)
}

func (uc renderPageRepo) GetCartFromDB(userID int64) (model.CartData, error) {
	return tx_repo.SelectCartByUserID(uc.dbCli, userID)
}

func (uc renderPageRepo) GetProductData(productID int64) (model.ProductData, error) {
	return product_repo.GetProductByProductID(uc.userCli, productID)
}

func (uc renderPageRepo) GetPromotion(coupon string, totalPrice float64) (model.PromotionData, error) {
	return promo_repo.GetPromotionByCoupon(uc.promotionCli, coupon, totalPrice)
}
