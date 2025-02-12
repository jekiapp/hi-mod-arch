package main

import (
	"log"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/internal/config"
	"github.com/jekiapp/hi-mod-arch/internal/domain"
	"github.com/jekiapp/hi-mod-arch/internal/logic"
	"github.com/jekiapp/hi-mod-arch/internal/usecase/post_payment"
	"github.com/jekiapp/hi-mod-arch/pkg/db"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
)

func initApplication(cfg *config.Config) Handler {

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
	promoCli := &http.Client{}

	createOrderUC := post_payment.NewCreateOrderUsecase(dbCli, productCli, promoCli)

	return Handler{
		cfg:                cfg,
		CreateOrderUsecase: createOrderUC,
	}
}

type Handler struct {
	cfg                *config.Config
	CreateOrderUsecase handler.GenericHandlerNsq
}

func (h *Handler) registerConsumer() []handler.Consumer {
	consumers := []handler.Consumer{}
	consumers = append(consumers,
		handler.NewGenericConsumer("mytopic", "channel_name", h.cfg.NsqConfig.ConsumerConfig, h.CreateOrderUsecase),
		// add more consumers here
	)
	return consumers
}
