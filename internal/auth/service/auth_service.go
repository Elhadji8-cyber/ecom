package service

import (
	"errors"
	"server/internal/auth/dto"
	"server/internal/auth/models"
	"server/internal/auth/repository"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
}

type authService struct {
	repo            repository.AuthRepository
	passwordService PasswordService
	jwtService      JwtService
}

func NewAuthService(repo repository.AuthRepository, pwdSvc PasswordService, jwtSvc JwtService) AuthService {
	return &authService{
		repo:            repo,
		passwordService: pwdSvc,
		jwtService:      jwtSvc,
	}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := s.passwordService.HashPassword(req.Password)
	if err != nil {
		return err
	}

	customer := &models.Customer{
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		Role:         "customer",
	}

	return s.repo.CreateCustomer(customer)
}

func (s *authService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	customer, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = s.passwordService.ComparePassword(req.Password, customer.PasswordHash)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtService.GenerateToken(customer.ID, customer.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: token}, nil
}
