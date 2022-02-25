package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
)

func SetQuant(router *gin.RouterGroup, handler *handler.QuantHandler) {
	quant := router.Group("/quants")
	{
		quant.GET("", handler.GetAllQuants)
		quant.GET("/quant", handler.GetQuant)

		quant.POST("/quant", handler.CreateQuant)

		quant.PATCH("/quant", handler.UpdateQuant)
		quant.PUT("/quant/option", handler.UpdateQuantOption)

		quant.DELETE("/quant", handler.DeleteQuant)
	}
}
