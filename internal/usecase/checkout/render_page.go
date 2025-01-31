package checkout

import (
	"database/sql"
	"hi-mod/internal/model"
	"hi-mod/pkg/httputil"
	"net/http"
)

type renderPageUsecase struct {
	dbCli        *sql.DB
	productCli   *http.Client
	promotionCli *http.Client
}

func RenderCheckoutPage(dbCli *sql.DB, promotionCli, productCli *http.Client) httputil.HttpHandler {
	return renderPageUsecase{
		dbCli:        dbCli,
		productCli:   productCli,
		promotionCli: promotionCli,
	}.renderPage
}

func (uc renderPageUsecase) renderPage(input interface{}) (output interface{}, err error) {

	finalPrice := logic.GetFinalPrice()

	response := model.CheckoutPageResponse{}
	return response, nil
}

func (uc renderPageUsecase) GetCartData() {
	//TODO implement me
	panic("implement me")
}

func (uc renderPageUsecase) GetProductInfo() {
	//TODO implement me
	panic("implement me")
}

func (uc renderPageUsecase) GetPromotion() {
	//TODO implement me
	panic("implement me")
}
