package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int            `gorm:"default:0" json:"stock"`
	CategoryID  uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	ImageURL    string         `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
