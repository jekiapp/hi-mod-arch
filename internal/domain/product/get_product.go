package product

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/model"
)

func GetProductByProductID(cfg *config.Config, cli *http.Client, productID int64) (model.ProductData, error) {
	// request to upstream to get product data
	url := fmt.Sprintf("https://%s/%s", cfg.Product.Host, cfg.Product.GetProductPath)
	cli.Get(url)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.ProductData{}, nil
}
