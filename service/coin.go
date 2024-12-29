package service

import (
	"fmt"

	"github.com/1067rail/portto-meme-coin/dto"
	"github.com/1067rail/portto-meme-coin/models"
	"gorm.io/gorm"
)

type CoinService struct {
	db *gorm.DB
}

func NewCoinService(db *gorm.DB) *CoinService {
	return &CoinService{
		db,
	}
}

func (c *CoinService) CreateCoin(payload dto.CreateCoinReq) (dto.CreateCoinRes, error) {
	memeCoin := models.MemeCoin{
		Name:            payload.Name,
		Description:     payload.Description,
		PopularityScore: 0,
	}

	if tx := c.db.Create(&memeCoin); tx.Error != nil {
		return dto.CreateCoinRes{}, tx.Error
	} else {
		return dto.CreateCoinRes{
			ID: memeCoin.ID,
		}, nil
	}

}

func (*CoinService) GetCoin() {
	fmt.Println("CoinService.GetCoin")
}

func (*CoinService) UpdateCoin() {
	fmt.Println("CoinService.UpdateCoin")
}

func (*CoinService) DeleteCoin() {
	fmt.Println("CoinService.DeleteCoin")
}

func (*CoinService) PokeCoin() {
	fmt.Println("CoinService.PokeCoin")
}