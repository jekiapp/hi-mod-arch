package product

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/config"
	"github.com/jekiapp/hi-mod-arch/internal/model"
)

func Init(cfg *config.Config) error {
	productURL = fmt.Sprintf("https://%s/%s", cfg.Product.Host, cfg.Product.GetProductPath)
	return nil
}

var productURL string

func GetProductByProductID(cli *http.Client, productID int64) (model.ProductData, error) {
	// request to upstream to get product data
	cli.Get(productURL)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.ProductData{}, nil
}
