package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/akiyamart/restAPIGo"
	"github.com/akiyamart/restAPIGo/pkg/repository"
)

const salt = "gsdkgflopxzcv814543sdaf3622o9"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService { 
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) { 
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string { 
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}