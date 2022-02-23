package routes

import (
	"economicus/internal/api/handler"
	"economicus/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	router  *gin.RouterGroup
	handler *handler.UserHandler
	mid     *middleware.AuthMiddleware
}

func NewUserRoute(router *gin.RouterGroup, handler *handler.UserHandler, mid *middleware.AuthMiddleware) *UserRoute {
	return &UserRoute{
		router:  router,
		handler: handler,
		mid:     mid,
	}
}

func (r *UserRoute) Setup() {
	r.router.POST("/register", r.handler.Register)
	userRoute := r.router.Group("/users", r.mid.Authenticate())
	{
		userRoute.GET("", r.handler.GetAllUsers)
		userRoute.GET("/user", r.handler.GetUser)
		userRoute.DELETE("/user", r.handler.DeleteUser)

		userRoute.PATCH("/profile", r.handler.EditUserProfile)
		userRoute.PUT("/profile-image", r.handler.UploadUserProfileImage)

		userRoute.GET("/favorite-quants", r.handler.GetFavoriteQuants)
		userRoute.POST("/favorite-quants", r.handler.AddToFavoriteQuants)
		userRoute.DELETE("/favorite-quants", r.handler.DeleteFromFavoriteQuants)

		userRoute.GET("/followings", r.handler.GetFollowings)
		userRoute.GET("/followers", r.handler.GetFollowers)
		userRoute.DELETE("/followings", r.handler.UnfollowUser)
		userRoute.POST("/followings", r.handler.FollowUser)
	}
}
