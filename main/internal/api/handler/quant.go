package handler

import (
	"economicus/internal/api/service"
	"economicus/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuantHandler struct {
	service *service.QuantService
}

func NewQuantHandler(s *service.QuantService) *QuantHandler {
	return &QuantHandler{
		service: s,
	}
}

// GetAllQuants returns list of quant models
func (h *QuantHandler) GetAllQuants(ctx *gin.Context) {
	option := models.NewQueryOption()

	if err := ctx.BindQuery(option); err != nil {
		sendQueryBindingErrMsg(ctx, err.Error())
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	quants, err := h.service.GetAllQuants(user.ID, option)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

// GetFollowingsQuants returns quants of followings
func (h *QuantHandler) GetFollowingsQuants(ctx *gin.Context) {
	option := models.NewQueryOption()

	if err := ctx.BindQuery(option); err != nil {
		sendQueryBindingErrMsg(ctx, err.Error())
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	quants, err := h.service.GetFollowingsQuants(user.ID, option)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

func (h *QuantHandler) GetQuantData(ctx *gin.Context) {
	var data models.QuantRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// GetQuant returns a quant models
func (h *QuantHandler) GetQuant(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	quant, err := h.service.GetQuant(user.ID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, quant)
}

// CreateQuant creates a quant models
func (h *QuantHandler) CreateQuant(ctx *gin.Context) {
	var request models.QuantRequest

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	res, err := h.service.CreateQuant(user.ID, &request)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// UpdateQuant updates a quant models
func (h *QuantHandler) UpdateQuant(ctx *gin.Context) {
	request := map[string]interface{}{}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.UpdateQuant(user.ID, request)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateQuantOption updates the quant's option
func (h *QuantHandler) UpdateQuantOption(ctx *gin.Context) {
	var request models.QuantOption

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.UpdateQuantOption(user.ID, &request)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteQuant deletes a quant models
func (h *QuantHandler) DeleteQuant(ctx *gin.Context) {
	var request struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.DeleteQuant(user.ID, request.QuantID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
