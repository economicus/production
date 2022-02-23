package repository

import (
	"economicus/internal/api/token"
	"economicus/internal/models"
	"time"
)

type AuthRepositoryFactory interface {
	ValidateUserWithToken(accessToken string) (uint, error)
	AuthenticateWithLocal(email, password string) (*token.JwtToken, error)
	RefreshToken(refreshToken string) (*token.JwtToken, error)
	//AuthenticateWithNaver(ctx *gin.Context) (*token.JwtToken, error)
	//AuthenticateWithKakao(ctx *gin.Context) (*token.JwtToken, error)
	//AuthenticateWithGoogle(ctx *gin.Context) (*token.JwtToken, error)
}

type UserRepositoryFactory interface {
	GetUsers(option *models.QueryOption) (models.Users, error)
	GetUserByID(userID uint) (*models.User, error)
	CheckNickname(nickname string) error
	CreateUser(email, password, name string) (uint, error)
	CreateProfile(userID uint, nickname string, birth time.Time) error
	UpdateUserProfile(userID uint, profile map[string]interface{}) error
	UploadUserProfileImage(userID uint, filepath string) error
	DeleteUser(userID uint) error
	UpdatePassword(userID uint, newPassword string) error

	GetFollowings(userID uint) (models.Users, error)
	GetFollowers(userID uint) (models.Users, error)
	Follow(userID, followerID uint) error
	UnFollow(userID, followerID uint) error

	GetFavoriteQuants(userID uint) ([]*models.Quant, error)
	AddToFavoriteQuants(userID, quantID uint) error
	DeleteFromFavoriteQuants(userID, quantID uint) error
}

// QuantRepositoryFactory is an interface which methods are related to quant model
type QuantRepositoryFactory interface {
	GetAllQuants(userID uint, option *models.QueryOption) (models.Quants, error)
	GetFollowingsQuants(userID uint, option *models.QueryOption) (models.Quants, error)
	GetQuant(quantID uint) (*models.Quant, error)
	GetMyQuants(userID uint) (models.Quants, error)

	CheckModelName(name string) error
	CheckQuantPermission(userID, quantID uint) error

	CreateQuant(quant *models.Quant) (uint, error)
	CreateQuantOption(quantOption *models.QuantOption) error

	UpdateQuant(quantID uint, data map[string]interface{}) error
	UpdateQuantOption(option *models.QuantOption) error

	DeleteQuant(quantID uint) error
}

// CommentRepositoryFactory is an interface which methods are related to comment
type CommentRepositoryFactory interface {
	GetCommentsAndReplies(quantID uint) (models.Comments, error)
	GetComment(commentID uint) (*models.Comment, error)
	CreateComment(comment *models.Comment) error
	UpdateComment(commentID uint, content string) error
	DeleteComment(commentID uint) error
}

// ReplyRepositoryFactory is an interface which methods are related to reply
type ReplyRepositoryFactory interface {
	GetReply(replyID uint) (*models.Reply, error)
	CreateReply(reply *models.Reply) error
	UpdateReply(replyID uint, content string) error
	DeleteReply(replyID uint) error
}
