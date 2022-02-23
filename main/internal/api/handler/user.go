package handler

import (
	"economicus/internal/api/hateos"
	"economicus/internal/api/service"
	"economicus/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
	hateos  *hateos.Hateos
}

func NewUserHandler(s *service.UserService, h *hateos.Hateos) *UserHandler {
	return &UserHandler{
		service: s,
		hateos:  h,
	}
}

// Register creates a user with a profile
func (h *UserHandler) Register(ctx *gin.Context) {
	var data models.RegisterRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	if err := h.service.Register(&data); err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"links": []map[string]string{h.hateos.LinkToLogin()},
	})
}

// GetAllUsers returns all user
func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	option := models.NewQueryOption()

	if err := ctx.BindQuery(option); err != nil {
		sendQueryBindingErrMsg(ctx, err.Error())
		return
	}

	users, err := h.service.GetUsers(option)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUser returns a user with id
func (h *UserHandler) GetUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	userID, err := extractUserId(user, ctx.Query("user_id"))
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	fields := getFieldsFromContext(ctx)

	resp, err := h.service.GetUserDataWithFields(userID, fields)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = h.service.DeleteUser(user.ID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// EditUserProfile edit a user profile
func (h *UserHandler) EditUserProfile(ctx *gin.Context) {
	request := map[string]interface{}{}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		sendErrMsgWithCode(ctx, http.StatusBadRequest, fmt.Sprintf("error while parsing json: %v", err))
		return
	}

	err = h.service.UpdateProfile(user.ID, request)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UploadUserProfileImage edit user's profile image
func (h *UserHandler) UploadUserProfileImage(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	file, header, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		sendErrMsgWithCode(ctx, http.StatusBadRequest, fmt.Sprintf("error while getting profile_image: %s", err))
		return
	}
	defer file.Close()

	err = h.service.UploadProfileImage(user.ID, file, header)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFollowings returns list of followings
func (h *UserHandler) GetFollowings(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	followings, err := h.service.GetFollowings(user.ID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": followings,
	})
}

// GetFollowers returns list of followers
func (h *UserHandler) GetFollowers(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	followings, err := h.service.GetFollowers(user.ID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": followings,
	})
}

// FollowUser refreshes access token
func (h *UserHandler) FollowUser(ctx *gin.Context) {
	var data struct {
		FollowerID uint `json:"follower_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.Follow(user.ID, data.FollowerID)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UnfollowUser refreshes access token
func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
	var data struct {
		FollowingID uint `json:"following_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.UnFollow(user.ID, data.FollowingID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFavoriteQuants returns a favorite quant list of user
func (h *UserHandler) GetFavoriteQuants(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	quants, err := h.service.GetFavoriteQuants(user.ID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

// AddToFavoriteQuants add a quant to favorite list
func (h *UserHandler) AddToFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.AddToFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// DeleteFromFavoriteQuants add a quant to favorite list
func (h *UserHandler) DeleteFromFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErrMsg(ctx, err, "")
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonBindingErrMsg(ctx, err.Error())
		return
	}

	err = h.service.DeleteFromFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendInternalErrMsg(ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
