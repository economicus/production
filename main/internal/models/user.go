package models

import (
	"economicus/commons"
	"gorm.io/gorm"
	"time"
)

type Users []*User

type User struct {
	gorm.Model     `json:"-" structs:"-"`
	UserActive     bool      `gorm:"column:user_active;default:true;" json:"-"`
	AccessLevel    int       `gorm:"column:access_level;default:1" json:"-"`
	LastLogin      time.Time `gorm:"column:last_login;default:null" json:"-"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Email          string    `gorm:"column:email;not null;unique" json:"email"`
	Password       []byte    `gorm:"column:password;not null" json:"-"`
	AuthResource   string    `gorm:"column:auth_resource;default:'local'"`
	Profile        Profile   `gorm:"constraint:OnDelete:CASCADE;" json:"profile"`
	Quants         Quants    `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	FavoriteQuants []*Quant  `gorm:"many2many:user_favorite_quants;" json:"-"`
	Comments       Comments  `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Replies        Replies   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Followings     Users     `gorm:"many2many:followings;" json:"-"`
}

func NewUser(email, password, name string) *User {
	u := User{
		Email:        email,
		Password:     []byte(password),
		Name:         name,
		AuthResource: "local",
	}
	u.HashPassword()
	return &u
}

func (u *User) HashPassword() {
	hashed, err := commons.HashPassword(u.Password)
	if err == nil {
		u.Password = hashed
	} else {
		logger.Warning(err)
	}
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetOwnerID() uint {
	return u.GetID()
}

func (u *User) ConvToMap() map[string]interface{} {
	data := map[string]interface{}{
		"name":    u.Name,
		"email":   u.Email,
		"profile": u.Profile.ConvToMap(),
	}
	return data
}
