package service

import (
	"errors"
	"fmt"
	"server/internal/auth/dto"
	"server/internal/auth/models"
	"server/internal/auth/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
	ForgotPassword(email string) error
	RefreshToken(token string) (*dto.AuthResponse, error)
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

func (s *authService) ForgotPassword(email string) error {
	customer, err := s.repo.FindByEmail(email)
	if err != nil {
		return errors.New("if an account with that email exists, we have sent a reset link")
	}

	// In a real app, generate a reset token, save it, and send an email.
	fmt.Printf("Reset password requested for: %s (ID: %s)\n", customer.Email, customer.ID)

	return nil
}

func (s *authService) RefreshToken(tokenStr string) (*dto.AuthResponse, error) {
	token, err := s.jwtService.ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userIDStr, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid user id in token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, errors.New("invalid user id format")
	}

	newToken, err := s.jwtService.GenerateToken(userID, role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: newToken}, nil
}
