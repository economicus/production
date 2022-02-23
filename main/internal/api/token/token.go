package token

import (
	"economicus/commons/converter"
	"economicus/config"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtToken(id uint, c *config.JwtConfig) (*JwtToken, error) {
	access, err := createToken(id, c, "access")
	if err != nil {
		return nil, err
	}
	refresh, err := createToken(id, c, "refresh")
	if err != nil {
		return nil, err
	}
	return &JwtToken{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

type Manager interface {
	Create(id uint) (*JwtToken, error)
	Validate(signedToken string) (uint, error)
}

type JwtManager struct {
	tokenConfig *config.JwtConfig
}

func NewJwtTokenManager(c *config.JwtConfig) *JwtManager {
	return &JwtManager{
		tokenConfig: c,
	}
}

func (m *JwtManager) Create(id uint) (*JwtToken, error) {
	return NewJwtToken(id, m.tokenConfig)
}

func (m *JwtManager) Validate(signedToken string) (uint, error) {
	var claims UserClaims
	token, err := jwt.ParseWithClaims(signedToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return []byte(m.tokenConfig.GetSecret()), nil
	})
	if err != nil {
		return 0, fmt.Errorf("error while parsing claims: %w", err)
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return converter.StrToUint(claims.Id)
}

func createToken(id uint, c *config.JwtConfig, subject string) (string, error) {
	var claims *UserClaims
	if subject == "access" {
		claims = NewUserClaims(id, c.AccessDuration, "access")
	} else {
		claims = NewUserClaims(id, c.RefreshDuration, "refresh")
	}
	token := jwt.NewWithClaims(c.Alg, claims)
	signedToken, err := token.SignedString([]byte(c.GetSecret()))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type UserClaims struct {
	jwt.StandardClaims
}

func NewUserClaims(userID uint, duration time.Duration, subject string) *UserClaims {
	var claims UserClaims
	claims.ExpiresAt = time.Now().Add(duration).Unix()
	claims.Id = fmt.Sprintf("%d", userID)
	claims.Subject = subject
	return &claims
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.Id == "0" {
		return fmt.Errorf("invalid user id")
	}
	return nil
}
