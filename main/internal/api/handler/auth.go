package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	e "main/internal/core/error"
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

// LoginInLocal logs in user
func (h *AuthHandler) LoginInLocal(ctx *gin.Context) {
	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}
	if req.Email == "" || req.Password == "" {
		sendErr(ctx, e.ErrMissingRequest)
		return
	}

	token, err := h.service.LoginInLocal(req.Email, req.Password)
	if err != nil {
		sendErr(ctx, err)
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
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	token, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
