package dto

import "time"

type PostRequestDTO struct {
	UserID    uint    `json:"user_id" binding:"required"`
	ProductID uint    `json:"product_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type PostResponseDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
