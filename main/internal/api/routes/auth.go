package routes

import (
	"economicus/internal/api/handler"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	router  *gin.RouterGroup
	handler *handler.AuthHandler
}

func NewAuthRoute(router *gin.RouterGroup, handler *handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		router:  router,
		handler: handler,
	}
}

func (a *AuthRoute) Setup() {
	a.router.POST("login", a.handler.Login)
	a.router.DELETE("/logout")
	a.router.POST("/refresh-token")
}
