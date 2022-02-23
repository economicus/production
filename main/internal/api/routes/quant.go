package routes

import (
	"economicus/internal/api/handler"
	"github.com/gin-gonic/gin"
)

type QuantRoute struct {
	router  *gin.RouterGroup
	handler *handler.QuantHandler
}

func NewQuantRoute(router *gin.RouterGroup, handler *handler.QuantHandler) *QuantRoute {
	return &QuantRoute{
		router:  router,
		handler: handler,
	}
}

func (q *QuantRoute) Setup() {
	quant := q.router.Group("/quants")
	{
		quant.GET("", q.handler.GetAllQuants)
		quant.GET("/quant", q.handler.GetQuant)
		quant.GET("/lab-data", q.handler.GetQuantData)

		quant.POST("/quant", q.handler.CreateQuant)

		quant.PATCH("/quant", q.handler.UpdateQuant)
		quant.PUT("/quant/option", q.handler.UpdateQuantOption)

		quant.DELETE("/quant", q.handler.DeleteQuant)
	}
}
