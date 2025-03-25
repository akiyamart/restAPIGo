package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	todo "github.com/akiyamart/restAPIGo"
	"github.com/akiyamart/restAPIGo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt 	   = "gsdkgflopxzcv814543sdaf3622o9"
	signingKey = "zxcvbnmasdfghjklqwewqfdsacxzf"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func generatePasswordHash(password string) string { 
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthService(repo repository.Authorization) *AuthService { 
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) { 
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func(s *AuthService) GenerateToken(username, password string) (string, error) { 
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil { 
		return "", err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) { 
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { 
			return nil, errors.New(("invalid signature method"))
		}
		
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok { 
		return 0, errors.New("token claims are not valid")
	}

	return claims.UserId, nil
}
