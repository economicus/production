package routes

import (
	"economicus/internal/api/handler"
	"github.com/gin-gonic/gin"
)

type CommentRoute struct {
	router  *gin.RouterGroup
	handler *handler.CommentHandler
}

func NewCommentRoute(router *gin.RouterGroup, handler *handler.CommentHandler) *CommentRoute {
	return &CommentRoute{
		router:  router,
		handler: handler,
	}
}

func (r *CommentRoute) Setup() {
	comment := r.router.Group("/comments")
	{
		comment.GET("", r.handler.GetCommentsAndReplies)
		comment.POST("", r.handler.CommentToQuant)
		comment.DELETE("", r.handler.DeleteComment)
		comment.PATCH("", r.handler.EditComment)
	}
}
