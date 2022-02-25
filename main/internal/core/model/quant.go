package model

import (
	"gorm.io/gorm"
)

type Quants []Quant

type Quant struct {
	gorm.Model  `json:"-"`
	UserID      uint        `json:"user_id"`
	QuantOption QuantOption `gorm:"constraint:OnDelete:CASCADE;foreignKey:QuantID;references:ID" json:"-"`
	Active      bool        `gorm:"column:active;default:false" json:"-"`
	LikedUsers  []*User     `gorm:"many2many:user_favorite_quants;" json:"-"`
	Name        string      `gorm:"column:name;not null;unique" json:"name"`
	Description string      `gorm:"column:description" json:"description"`
	Private     bool        `gorm:"column:private;default:false" json:"-"`
	Comments    Comments    `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}

func NewQuant(userID uint, name string) *Quant {
	return &Quant{
		UserID: userID,
		Name:   name,
	}
}

func (q *Quant) GetID() uint {
	return q.ID
}

func (q *Quant) GetOwnerID() uint {
	return q.UserID
}

func (q *Quant) ToMap() map[string]interface{} {
	return ToMap(q)
}

func (q *Quant) ToMapWithFields(fields []string) map[string]interface{} {
	return toMapWithFields(q, fields)
}
