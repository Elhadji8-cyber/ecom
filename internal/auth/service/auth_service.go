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

// AuthService defines the interface for authentication-related operations.
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
	existen, _ := s.repo.FindByEmail(req.Email)
	if existen != nil {
		return errors.New("email already existed")
	}
	hashedPassword, err := s.passwordService.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
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
	err = s.passwordService.ComparePasswords(req.Password, customer.PasswordHash)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := s.jwtService.GenerateToken(customer.ID, customer.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}
	return &dto.AuthResponse{
		Token: token,
	}, nil
}

func (s *authService) ForgotPassword(email string) error {
	//--------------------------- Implement forgot password logic ------------------//
	customer, err := s.repo.FindByEmail(email)
	if err != nil {
		return errors.New("if an account with email exists, we have send a reset link")
	}
	// Generate reset token and send email logic here
	fmt.Printf("Reset password link sent to %s\n", customer.Email)
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
		return nil, errors.New("invalid user id Token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role in token")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, errors.New("invalid user id format in token")
	}
	newToken, err := s.jwtService.GenerateToken(userID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new token: %v", err)
	}
	return &dto.AuthResponse{Token: newToken}, nil
}
