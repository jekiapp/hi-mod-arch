package checkout

type renderPageInterface interface {
	GetCartData()
	GetProductInfo()
	GetFinalPrice()
}
