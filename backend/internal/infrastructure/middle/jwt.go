package middle

import (
	"time"

	"github.com/KenethSandoval/doc-md/internal/domain"
	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	UserJWT struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	JWTClaim struct {
		jwtlib.StandardClaims
		SessID         int         `json:"sess_id"`
		User           UserJWT     `json:"user"`
		Lang           interface{} `json:"lang"`
		SessionSetting int         `json:"session_setting"`
	}
)

func JWT(key string, wl ...string) echo.MiddlewareFunc {
	jwtCfg := middleware.JWTConfig{
		Claims:     &JWTClaim{},
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return inArrayString(wl, path)
		},
	}
	return middleware.JWTWithConfig(jwtCfg)
}

func CreateJWTToken(key string, payload *domain.User) (string, error) {
	claims := &JWTClaim{
		SessID: time.Now().Nanosecond(),
		User: UserJWT{
			ID:       1,
			Username: payload.Username,
			Role:     payload.Role,
		},
		Lang:           nil,
		SessionSetting: 1,
		StandardClaims: jwtlib.StandardClaims{
			Audience:  "audien",
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			Id:        "id",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "issuer",
			NotBefore: 0,
			Subject:   "subject",
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func inArrayString(arr []string, s ...string) bool {
	if len(arr) == 0 || len(s) == 0 {
		return false
	}
	for _, ar := range arr {
		for _, v := range s {
			if ar == v {
				return true
			}
		}
	}
	return false
}
