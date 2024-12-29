package dto

import "github.com/google/uuid"

type CreateCoinReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCoinRes struct {
	ID uuid.UUID `json:"id"`
}
