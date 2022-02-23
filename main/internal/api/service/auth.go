package service

import (
	"economicus/internal/api/repository"
	"economicus/internal/api/token"
)

type AuthService struct {
	repo repository.AuthRepositoryFactory
}

func NewAuthService(repo repository.AuthRepositoryFactory) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(email, password string) (*token.JwtToken, error) {
	return s.repo.AuthenticateWithLocal(email, password)
}

func (s *AuthService) RefreshToken(refreshToken string) (*token.JwtToken, error) {
	return s.repo.RefreshToken(refreshToken)
}
