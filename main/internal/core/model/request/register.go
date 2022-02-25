package request

import (
	"time"
)

type RegisterRequest struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Birth    time.Time `json:"birth"`
}
