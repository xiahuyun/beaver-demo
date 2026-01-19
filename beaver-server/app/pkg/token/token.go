package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClient struct {
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

type JwtPayLoad struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func (t *TokenClient) GenerateToken(payload JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expires))),
			Issuer:    "usercenter",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func (t *TokenClient) VerifyToken(tokenString string, accessSecret string) (JwtPayLoad, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return JwtPayLoad{}, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return JwtPayLoad{}, errors.New("token 无效")
	}
	return claims.JwtPayLoad, nil
}
