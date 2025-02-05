package checkout

import (
	"database/sql"
	"github.com/jekiapp/hi-mod/internal/logic"
	"github.com/jekiapp/hi-mod/internal/model"
	"github.com/jekiapp/hi-mod/pkg/httputil"
	"net/http"
)

type renderPageUsecase struct {
	dbCli        *sql.DB
	productCli   *http.Client
	promotionCli *http.Client
}

func RenderCheckoutPage(dbCli *sql.DB, promotionCli, productCli *http.Client) httputil.GenericHandler {
	return renderPageUsecase{
		dbCli:        dbCli,
		productCli:   productCli,
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

	}
	err = logic.CalculatePrice(&cartData, uc)
	if err != nil {

	}

	response := model.CheckoutPageResponse{}
	return response, nil
}

func (uc renderPageUsecase) GetCartFromDB(userID int64) (model.CartData, error) {

}

func (uc renderPageUsecase) GetProductPrice(productID int64) (float64, error) {

}

func (uc renderPageUsecase) GetPromotion(userID, productID int64) (model.PromotionData, error) {

}
