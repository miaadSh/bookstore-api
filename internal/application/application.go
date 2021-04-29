package application

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/pkg/translator/i18n"
)

func Run(cfg *config.Config) error {

	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	_ = translator

	return err
}