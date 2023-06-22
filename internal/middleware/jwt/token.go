package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"vax/internal/config"
)

func GenerateToken(uid int64, name string) (string, error) {
	claims := jwt.StandardClaims{
		// 接受者
		Audience: name,
		// 过期时间
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		// 唯一表示符
		Id: strconv.FormatInt(uid, 10),
		// 签发时间
		IssuedAt: time.Now().Unix(),
		// 签发者
		Issuer: "AboutVaccine",
		// 生效时间
		NotBefore: time.Now().Unix(),
		// 主题
		Subject: "token",
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString([]byte(config.JwtSecret)); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	// 解析JWT Token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.JwtSecret), nil
	})

	// 如果解析成功
	if err == nil && token != nil {
		claim, ok := token.Claims.(*jwt.StandardClaims)
		if ok && token.Valid {
			return claim, nil
		}
	}
	return nil, err
}
