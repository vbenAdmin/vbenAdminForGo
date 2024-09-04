package Jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(userId uint, secretKey string) (string, int64, error) {
	expAt := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	//expAt := jwt.NewNumericDate(time.Unix(int64(1625537464), 0))

	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: expAt,                          // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secretKey))
	return token, expAt.Unix(), err
}

// ParseJwt 解析JWT
func ParseJwt(tokenString, secretKey string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
