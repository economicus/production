package handler

import (
	"economicus/internal/error"
	"economicus/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func getUserFromContext(ctx *gin.Context) (*models.User, error) {
	setUser, exist := ctx.Get("user")
	if !exist {
		return nil, fmt.Errorf("user is not exists in context")
	}
	user, ok := setUser.(models.User)
	if !ok {
		return nil, fmt.Errorf("user's type is invalid")
	}
	return &user, nil
}

func getFieldsFromContext(ctx *gin.Context) []string {
	field := ctx.Query("fields")
	if field == "" {
		return []string{}
	}
	return strings.Split(field, ",")
}

func extractUserId(user *models.User, query string) (uint, error) {
	if query == "" {
		return user.ID, nil
	}
	userID64, err := strconv.ParseUint(query, 10, 64)
	if err != nil {
		return 0, ecoerror.ErrTypeConvertFailed
	}
	return uint(userID64), nil
}
