package controller

import (
	"github.com/1067rail/portto-meme-coin/service"
	"gorm.io/gorm"
)

type Controller struct {
	coinService *service.CoinService
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		coinService: service.NewCoinService(db),
	}
}
