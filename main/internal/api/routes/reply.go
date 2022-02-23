package routes

import (
	"economicus/internal/api/handler"
	"github.com/gin-gonic/gin"
)

type ReplyRoute struct {
	router  *gin.RouterGroup
	handler *handler.ReplyHandler
}

func NewReplyRoute(router *gin.RouterGroup, handler *handler.ReplyHandler) *ReplyRoute {
	return &ReplyRoute{
		router:  router,
		handler: handler,
	}
}

func (r *ReplyRoute) Setup() {
	reply := r.router.Group("/replies")
	{
		reply.POST("", r.handler.ReplyToComment)
		reply.DELETE("", r.handler.DeleteReply)
		reply.PATCH("", r.handler.EditReply)
	}
}
