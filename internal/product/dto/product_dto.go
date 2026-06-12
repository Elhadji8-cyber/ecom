package dto

import "github.com/google/uuid"

type ProductRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	Stock       int       `json:"stock" binding:"required,gte=0"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required"`
	ImageURL    string    `json:"image_url"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
	CategoryName string   `json:"category_name"`
	ImageURL    string    `json:"image_url"`
}
