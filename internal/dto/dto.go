package dto

import "time"

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
