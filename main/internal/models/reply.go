package models

import (
	"gorm.io/gorm"
)

type Replies []Reply

type Reply struct {
	gorm.Model `json:"-"`
	CommentID  uint   `json:"comment_id"`
	UserID     uint   `json:"user_id"`
	Content    string `gorm:"type:text;column:content" json:"content"`
}

func NewReply(userID, commentID uint, content string) *Reply {
	return &Reply{
		UserID:    userID,
		CommentID: commentID,
		Content:   content,
	}
}

func (r *Reply) GetID() uint {
	return r.ID
}

func (r *Reply) GetOwnerID() uint {
	return r.UserID
}

func (r *Reply) ConvToMap() map[string]interface{} {
	m := map[string]interface{}{}
	m["user_id"] = r.UserID
	m["quant_id"] = r.CommentID
	m["content"] = r.Content
	return m
}
