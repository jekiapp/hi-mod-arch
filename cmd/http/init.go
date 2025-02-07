package main

import (
	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/usecase/checkout"
	"github.com/jekiapp/hi-mod/pkg/db"
	"log"
	"net/http"
)

func InitApplication() Handler {
	cfg := config.InitConfig()
	// init database
	dbCli, err := db.InitDatabase(db.DbConfig{Host: cfg.Database.Host})
	if err != nil {
		log.Fatal(err)
	}

	productCli := &http.Client{}
	userCli := &http.Client{}
	promoCli := &http.Client{}

	newHandler := Handler{
		CheckoutPageHandler: checkout.RenderCheckoutPage(cfg, dbCli, promoCli, productCli, userCli),
	}

	return newHandler
}
