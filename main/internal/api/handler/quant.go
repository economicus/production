package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	"main/internal/core/model"
	"main/internal/core/model/request"
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

// GetAllQuants returns list of quant model
func (h *QuantHandler) GetAllQuants(ctx *gin.Context) {
	option := model.NewQuery()

	if err := ctx.BindQuery(option); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	quants, err := h.service.GetAllQuants(user.ID, option)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

// GetQuant returns a quant models
func (h *QuantHandler) GetQuant(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	quant, err := h.service.GetQuant(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, quant)
}

// CreateQuant creates a quant model
func (h *QuantHandler) CreateQuant(ctx *gin.Context) {
	var req request.QuantC

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	res, err := h.service.CreateQuant(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// UpdateQuant updates a quant model
func (h *QuantHandler) UpdateQuant(ctx *gin.Context) {
	var req request.QuantE

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.UpdateQuant(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateQuantOption updates the quant's option
func (h *QuantHandler) UpdateQuantOption(ctx *gin.Context) {
	var req model.QuantOption

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.UpdateQuantOption(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteQuant deletes a quant model
func (h *QuantHandler) DeleteQuant(ctx *gin.Context) {
	var req struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.DeleteQuant(user.ID, req.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
