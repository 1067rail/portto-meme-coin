package database

import (
	"fmt"

	"github.com/1067rail/portto-meme-coin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	appConfig := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		appConfig.DatabaseHost,
		appConfig.DatabaseUser,
		appConfig.DatabasePassword,
		appConfig.DatabaseDbName,
		appConfig.DatabasePort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
