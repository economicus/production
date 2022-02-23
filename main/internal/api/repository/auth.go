package repository

import (
	"economicus/commons"
	"economicus/internal/api/token"
	"economicus/internal/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type AuthRepository struct {
	tokenManager token.Manager
	db           *gorm.DB
	logger       *log.Logger
}

func NewAuthRepository(db *gorm.DB, tokenManager token.Manager, logger *log.Logger) AuthRepositoryFactory {
	return &AuthRepository{
		tokenManager: tokenManager,
		db:           db,
		logger:       logger,
	}
}

func (repo *AuthRepository) ValidateUserWithToken(accessToken string) (uint, error) {
	userID, err := repo.tokenManager.Validate(accessToken)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (repo *AuthRepository) RefreshToken(refreshToken string) (*token.JwtToken, error) {
	userID, err := repo.tokenManager.Validate(refreshToken)
	if err != nil {
		return nil, err
	}
	return repo.tokenManager.Create(userID)
}

// AuthenticateWithLocal authenticates user with email and password which stored in local database
func (repo *AuthRepository) AuthenticateWithLocal(email, password string) (*token.JwtToken, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).First(&user).Update("last_login", time.Now())
	if result.Error != nil {
		return nil, result.Error
	}
	err := commons.ComparePassword([]byte(password), user.Password)
	if err != nil {
		return nil, err
	}
	tokens, err := repo.tokenManager.Create(user.ID)
	if err != nil {
		repo.logger.Errorf("error while creating token: %v", err)
		return nil, fmt.Errorf("error in Authenticate while creating token: %w", err)
	}
	return tokens, nil
}
