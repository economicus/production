package config

import (
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

type JwtConfig struct {
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	Alg             jwt.SigningMethod

	secret string
}

func NewJwtConfig() *JwtConfig {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatalf("error while getting secret key: empty secret")
	}
	return &JwtConfig{
		AccessDuration:  time.Hour,
		RefreshDuration: time.Hour * 24 * 7,
		Alg:             jwt.SigningMethodHS512,
		secret:          secret,
	}
}

func (c *JwtConfig) GetSecret() string {
	return c.secret
}
