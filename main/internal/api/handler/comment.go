package handler

import (
	"economicus/internal/api/hateos"
	"economicus/internal/api/service"
	"economicus/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	service *service.CommentService
	hateos  *hateos.Hateos
}

func NewCommentHandler(s *service.CommentService, h *hateos.Hateos) *CommentHandler {
	return &CommentHandler{
		service: s,
		hateos:  h,
	}
}

// GetCommentsAndReplies returns all comments and replies of a quant models
func (h *CommentHandler) GetCommentsAndReplies(ctx *gin.Context) {
	quantIdStr := ctx.Query("quant_id")
	if quantIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting quant id",
		})
		return
	}
	quantID, err := strconv.ParseUint(quantIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("error while converting quant id: %s", err),
		})
		return
	}
	comments, err := h.service.GetCommentsAndReplies(uint(quantID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("error while getting comments: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

// CommentToQuant create a comment
func (h *CommentHandler) CommentToQuant(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while getting user from request: %s", err),
		})
		return
	}
	var data struct {
		QuantID uint   `json:"quant_id"`
		Content string `json:"content"`
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error while parsing json: %s", err),
		})
		return
	}
	err = h.service.CreateComment(user.ID, data.QuantID, data.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while creating comment"),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

// DeleteComment delete a comment
func (h *CommentHandler) DeleteComment(ctx *gin.Context) {
	object, exist := ctx.Get("object")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting comment",
		})
		return
	}
	comment, _ := object.(*models.Comment)
	err := h.service.DeleteComment(comment.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while deleting comment: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// EditComment edit a comment
func (h *CommentHandler) EditComment(ctx *gin.Context) {
	object, exist := ctx.Get("object")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting comment",
		})
		return
	}
	comment, _ := object.(*models.Comment)
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
		return
	}
	err := h.service.UpdateComment(comment.ID, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error while updating comment: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
