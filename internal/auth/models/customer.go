package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id,omitempty"`
	FullName     string         `gorm:"type:varchar(255);not null" json:"full_name"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone        string         `gorm:"type:varchar(20)" json:"phone"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	Role         string         `gorm:"type:varchar(50);default:'customer'" json:"role"`
	IsVerified   bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
