package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemeCoin struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name            string    `gorm:"type:varchar(50);unique;not null"`
	Description     string    `grom:"type:ntext"`
	PopularityScore int       `gorm:"type:bigint"`
}
