package repository

import (
	"server/internal/auth/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateCustomer(customer *models.Customer) error
	FindByEmail(email string) (*models.Customer, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *authRepository) FindByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
