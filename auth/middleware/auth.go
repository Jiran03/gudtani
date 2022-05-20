package authMiddleware

import (
	"time"

	"github.com/Jiran03/gudtani/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (cJWT ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(auth.SECRET_KEY),
	}
}

//Generate Token
func (cJWT *ConfigJWT) GenerateToken(userID int) string {
	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cJWT.ExpiresDuration))).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(auth.SECRET_KEY))
	return token
}

//Get User from JWT
func GetUser(ctx echo.Context) *JWTCustomClaims {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	return claims
}
