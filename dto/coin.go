package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCoinReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCoinRes struct {
	ID uuid.UUID `json:"id"`
}

type GetCoinRes struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore int       `json:"popularity_score"`
}
