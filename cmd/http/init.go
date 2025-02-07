package main

import (
	"log"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/domain"
	"github.com/jekiapp/hi-mod/internal/logic"
	"github.com/jekiapp/hi-mod/internal/usecase/checkout"
	"github.com/jekiapp/hi-mod/pkg/db"
)

func InitApplication() Handler {
	cfg := config.InitConfig()

	err := logic.Init(cfg)
	if err != nil {
		log.Fatalf("error init logic %s", err.Error())
	}

	err = domain.Init(cfg)
	if err != nil {
		log.Fatalf("error init domain %s", err.Error())
	}

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
