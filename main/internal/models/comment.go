package models

import (
	"economicus/commons/converter"
	"gorm.io/gorm"
)

type Comments []Comment

type Comment struct {
	gorm.Model `json:"-" structs:"-"`
	UserID     uint    `json:"user_id"`
	QuantID    uint    `json:"quant_id"`
	Content    string  `gorm:"type:text;column:content" json:"content"`
	Replies    Replies `gorm:"constraint:OnDelete:CASCADE;" json:"replies"`
}

func NewComment(userID, quantID uint, content string) *Comment {
	return &Comment{
		UserID:  userID,
		QuantID: quantID,
		Content: content,
	}
}

func (c *Comment) GetID() uint {
	return c.ID
}

func (c *Comment) GetOwnerID() uint {
	return c.UserID
}

func (c *Comment) ConvToMap() map[string]interface{} {
	return converter.InterfaceToMap(c)
}
