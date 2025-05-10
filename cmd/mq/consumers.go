package main

import (
	"log"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/config"
	"github.com/jekiapp/hi-mod-arch/internal/logic"
	"github.com/jekiapp/hi-mod-arch/internal/model"
	"github.com/jekiapp/hi-mod-arch/internal/repository"
	"github.com/jekiapp/hi-mod-arch/internal/usecase/post_payment"
	"github.com/jekiapp/hi-mod-arch/pkg/db"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
)

func initApplication(cfg *config.Config) Handler {

	err := logic.Init(cfg)
	if err != nil {
		log.Fatalf("error init logic %s", err.Error())
	}

	err = repository.Init(cfg)
	if err != nil {
		log.Fatalf("error init repository %s", err.Error())
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
	CreateOrderUsecase handler.GenericHandlerNsq[model.PaymentSuccess]
}

func (h *Handler) registerConsumer() []handler.Consumer {
	consumers := []handler.Consumer{}
	consumers = append(consumers,
		handler.NewGenericConsumer("mytopic", "channel_name", h.cfg.NsqConfig.ConsumerConfig, h.CreateOrderUsecase),
		// add more consumers here
	)
	return consumers
}
