package handler

import (
	"economicus/internal/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendErrMsgWithCode(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendJsonBindingErrMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("error while parsing json: %s", msg),
	})
}

func sendQueryBindingErrMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("error while parsing query: %s", msg),
	})
}

// sendErrMsg sends a message with different errors
func sendErrMsg(ctx *gin.Context, err error, msg string) {
	var code int
	fullMsg := err.Error()
	if msg != "" {
		fullMsg += ": " + msg
	}
	switch err {
	case ecoerror.ErrDuplicateEmail,
		ecoerror.ErrDuplicateNickname,
		ecoerror.ErrInvalidPassword,
		ecoerror.ErrDuplicateModelName,
		ecoerror.ErrTypeConvertFailed:
		code = http.StatusBadRequest
	case ecoerror.ErrInactiveAccount:
		code = http.StatusNotFound
		fullMsg += ": " + msg
	case ecoerror.ErrNoAuthorization,
		ecoerror.ErrPermissionDenied:
		code = http.StatusForbidden
	case ecoerror.ErrNoRecord:
		code = http.StatusBadRequest
	default:
		sendInternalErrMsg(ctx)
		return
	}
	ctx.JSON(code, gin.H{
		"message": fullMsg,
	})
}

// sendInternalErrMsg sends a message with http status code 500
func sendInternalErrMsg(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "something went wrong! please contact to manager",
	})
}
