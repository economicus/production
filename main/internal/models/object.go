package models

type Object interface {
	GetID() uint
	GetOwnerID() uint
	ConvToMap() map[string]interface{}
}
