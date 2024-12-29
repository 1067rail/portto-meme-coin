package service

import (
	"fmt"

	"github.com/1067rail/portto-meme-coin/dto"
	"github.com/1067rail/portto-meme-coin/models"
	"github.com/google/uuid"
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

func (c *CoinService) GetCoin(id uuid.UUID) (dto.GetCoinRes, error) {
	memeCoin := models.MemeCoin{
		ID: id,
	}

	if tx := c.db.First(&memeCoin); tx.Error != nil {
		return dto.GetCoinRes{}, tx.Error
	} else {
		return dto.GetCoinRes{
			Name:            memeCoin.Name,
			Description:     memeCoin.Description,
			CreatedAt:       memeCoin.CreatedAt,
			PopularityScore: memeCoin.PopularityScore,
		}, nil
	}
}

func (c *CoinService) UpdateCoin(id uuid.UUID, payload dto.UpdateCoinReq) (dto.UpdateCoinRes, error) {

	err := c.db.Transaction(func(db *gorm.DB) error {

		memeCoin := models.MemeCoin{
			ID: id,
		}

		if tx := c.db.First(&memeCoin); tx.Error != nil {
			return tx.Error
		}

		memeCoin.Description = payload.Description

		db.Save(memeCoin)

		return nil
	})

	if err != nil {
		return dto.UpdateCoinRes{}, err
	}

	return dto.UpdateCoinRes{ID: id}, nil
}

func (c *CoinService) DeleteCoin(id uuid.UUID) (dto.DeleteCoinRes, error) {
	if tx := c.db.Unscoped().Delete(&models.MemeCoin{}, id); tx.Error != nil {
		return dto.DeleteCoinRes{}, tx.Error
	}
	return dto.DeleteCoinRes{ID: id}, nil
}

func (*CoinService) PokeCoin() {
	fmt.Println("CoinService.PokeCoin")
}
