package handler

import (
	"economicus/internal/api/hateos"
	"economicus/internal/api/service"
	"economicus/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReplyHandler struct {
	service *service.ReplyService
	hateos  *hateos.Hateos
}

func NewReplyHandler(service *service.ReplyService, hateos *hateos.Hateos) *ReplyHandler {
	return &ReplyHandler{
		service: service,
		hateos:  hateos,
	}
}

// ReplyToComment create a reply to a comment
func (h *ReplyHandler) ReplyToComment(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while getting user from request: %s", err),
		})
		return
	}
	var data struct {
		CommentID uint   `json:"comment_id"`
		Content   string `json:"content"`
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while parsing json: %s", err),
		})
		return
	}
	err = h.service.CreateReply(user.ID, data.CommentID, data.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while creating reply"),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

// DeleteReply delete a reply
func (h *ReplyHandler) DeleteReply(ctx *gin.Context) {
	object, exist := ctx.Get("object")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting comment",
		})
		return
	}
	reply, _ := object.(*models.Reply)
	err := h.service.DeleteReply(reply.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while deleting reply: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// EditReply edit a reply
func (h *ReplyHandler) EditReply(ctx *gin.Context) {
	object, exist := ctx.Get("object")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting comment",
		})
		return
	}
	reply, _ := object.(*models.Reply)
	dataInterface, _ := ctx.Get("data")
	data, _ := dataInterface.(map[string]interface{})
	contentInterface, exist := data["content"]
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting content interface",
		})
		return
	}
	content, ok := contentInterface.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "content type err",
		})
	}
	err := h.service.UpdateReply(reply.ID, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while updating reply: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
