package main

import (
	"log"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/internal/config"
	"github.com/jekiapp/hi-mod-arch/internal/domain"
	"github.com/jekiapp/hi-mod-arch/internal/logic"
	"github.com/jekiapp/hi-mod-arch/internal/usecase/checkout"
	"github.com/jekiapp/hi-mod-arch/pkg/db"
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
		CheckoutPageHandler: checkout.RenderCheckoutPage(dbCli, promoCli, productCli, userCli),
	}

	return newHandler
}
