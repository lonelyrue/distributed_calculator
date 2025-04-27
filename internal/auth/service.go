package auth

import (
	"distributed_calculator/internal/storage"
	"errors"
)

type AuthService struct {
	storage *storage.Storage
}

func NewAuthService(storage *storage.Storage) *AuthService {
	return &AuthService{storage: storage}
}

func (a *AuthService) Register(login, password string) error {
	return a.storage.CreateUser(login, password)
}

func (a *AuthService) Login(login, password string) (string, error) {
	user, err := a.storage.GetUserByLogin(login)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("invalid credentials")
	}
	return GenerateJWT(user.ID)
}

func (a *AuthService) GetUserIDFromToken(tokenStr string) (int, error) {
	claims, err := ParseJWT(tokenStr)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
