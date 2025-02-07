package promo

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/model"
)

func GetPromotionByCoupon(cfg *config.Config, cli *http.Client, coupon string, totalPrice float64) (model.PromotionData, error) {
	// request to upstream to get promotion data
	url := fmt.Sprintf("https://%s/%s", cfg.Promo.Host, cfg.Promo.GetPromoPath)
	cli.Get(url)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.PromotionData{}, nil
}
