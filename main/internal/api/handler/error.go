package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	e "main/internal/core/error"
	"net/http"
)

func sendErr(ctx *gin.Context, err error) {
	code := e.DetErrCode(err)
	if code == http.StatusInternalServerError {
		sendInternalErr(ctx)
		return
	}
	ctx.JSON(code, gin.H{
		"message": err.Error(),
	})
}

func sendErrWithMsg(ctx *gin.Context, err error, msg string) {
	code := e.DetErrCode(err)
	if code == http.StatusInternalServerError {
		sendInternalErr(ctx)
		return
	}
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("%s: %s", err.Error(), msg),
	})
}

func sendJsonParsingErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("error while parsing json: %v", err),
	})
}

func sendInternalErr(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong. Please contact economicus members.",
	})
}
