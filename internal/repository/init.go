package repository

import (
	"github.com/jekiapp/hi-mod-arch/config"
	"github.com/jekiapp/hi-mod-arch/internal/repository/product"
	"github.com/jekiapp/hi-mod-arch/internal/repository/promo"
	"github.com/jekiapp/hi-mod-arch/internal/repository/user"
)

func Init(cfg *config.Config) error {
	var err error

	err = promo.Init(cfg)
	if err != nil {
		return err
	}

	err = user.Init(cfg)
	if err != nil {
		return err
	}

	err = product.Init(cfg)
	if err != nil {
		return err
	}

	// init other repository
	return err
}
