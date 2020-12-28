package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tarpal7/all-commerce"
	"github.com/tarpal7/all-commerce/pkg/repository"
)
const (
	salt = "defxcvxopwerwl345ds25gjmxsdf"
	signingKey="sdfsdfssd67823jlsdbxc()*&234"
	tokenTTL = 12* time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService (repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user allcommerce.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken ...
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
	},
	user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// ParseToken ...
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}