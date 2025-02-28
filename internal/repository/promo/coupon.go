package promo

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/config"
	"github.com/jekiapp/hi-mod-arch/internal/model"
)

func Init(cfg *config.Config) error {

	promoURL = fmt.Sprintf("https://%s/%s", cfg.Promo.Host, cfg.Promo.GetPromoPath)

	return nil
}

var promoURL string

func GetPromotionByCoupon(cli *http.Client, coupon string, totalPrice float64) (model.PromotionData, error) {
	// request to upstream to get promotion data
	cli.Get(promoURL)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.PromotionData{}, nil
}
