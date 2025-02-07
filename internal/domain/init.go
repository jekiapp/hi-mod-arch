package domain

import (
	"github.com/jekiapp/hi-mod-arch/internal/config"
	"github.com/jekiapp/hi-mod-arch/internal/domain/product"
	"github.com/jekiapp/hi-mod-arch/internal/domain/promo"
	"github.com/jekiapp/hi-mod-arch/internal/domain/user"
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

	// init other domain
	return err
}
