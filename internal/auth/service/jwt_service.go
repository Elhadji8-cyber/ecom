package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtService interface {
	GenerateToken(userID uuid.UUID, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey     string
	expirationHrs time.Duration
}

func NewJwtService() JwtService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("Warning: JWT_SECRET not set, using insecure default")
		secret = "insecure_default_secret"
	}

	expStr := os.Getenv("JWT_EXPIRATION")
	expHrs := 24
	if expStr != "" {
		if d, err := time.ParseDuration(expStr); err == nil {
			expHrs = int(d.Hours())
		}
	}

	return &jwtService{
		secretKey:     secret,
		expirationHrs: time.Duration(expHrs),
	}
}

func (s *jwtService) GenerateToken(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(time.Hour * s.expirationHrs).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
}
