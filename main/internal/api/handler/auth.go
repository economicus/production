package handler

import (
	"economicus/internal/api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

// Login logs in user
func (h *AuthHandler) Login(ctx *gin.Context) {
	input := struct {
		Email    string
		Password string
	}{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil && input.Email != "" && input.Password != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while parsing json: %s", err),
		})
		return
	}
	token, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("error while authenticating: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, token)
}

// Logout logs out user
func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "not set up yet",
	})
}

// RefreshToken refreshes access token
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var input map[string]string
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while parsing json: %s", err),
		})
		return
	}
	refreshToken := input["refresh_token"]
	if refreshToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while getting refreshToken: refreshToken is an empty string"),
		})
		return
	}
	accessToken, err := h.service.RefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("error while refreshing access token: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
