package middleware

import (
	"economicus/internal/api/token"
	"economicus/internal/drivers"
	"economicus/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	db  *drivers.DB
	jwt *token.JwtManager
}

func NewAuthMiddleware(db *drivers.DB, jwt *token.JwtManager) *AuthMiddleware {
	return &AuthMiddleware{
		db:  db,
		jwt: jwt,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := extractAccessToken(ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error while extracting access token: %s", err),
			})
			ctx.Abort()
			return
		}
		userID, err := m.jwt.Validate(accessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		user, err := getUserByID(m.db.SQL, userID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("error while finding a user: %s", err),
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", *user)
		ctx.Next()
	}
}

func getUserByID(db *gorm.DB, userID uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// extractAccessToken extracts access token from request
func extractAccessToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	data := strings.Split(header, " ")
	if len(data) != 2 {
		return "", fmt.Errorf("error in extractAccessToken while split header")
	}
	return data[1], nil
}
