package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pass dbname=meme-coin port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
