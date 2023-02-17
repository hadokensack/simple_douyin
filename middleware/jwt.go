package middleware

import (
	"simple_douyin/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("tiktok")

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

// ReleaseToken 颁发token
func ReleaseToken(user models.UserLogin) (string, error) {
	expirationTime := time.Now().Add(7 * time.Hour)
	claims := &Claims{
		UserId: user.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //失效时间
			IssuedAt:  time.Now().Unix(),     //签发时间
			Issuer:    "Zhao",                //签发人
			Subject:   "token",               //主题
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}
