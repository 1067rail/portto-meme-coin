package service

import (
	"fmt"

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

func (*CoinService) CreateCoin() {
	fmt.Println("CoinService.CreateCoin")
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
