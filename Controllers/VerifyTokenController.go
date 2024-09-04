package Controllers

import (
	"github.com/gin-gonic/gin"
	"goVben/Redis"
	"goVben/Utils/Config"
	"goVben/Utils/Jwt"
	"goVben/Utils/Resful"
	"net/http"
	"time"
)

func VerifyTokenController(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "未登录", nil))
		return
	}
	claims, err := Jwt.ParseJwt(authHeader, Config.Conf.JwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "token非法", err.Error()))
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "token失效", nil))
		return
	}
	token, err := Redis.GetUserToken(claims.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "token非法了", nil))
		return
	}
	if claims.ExpiresAt.Unix() != token {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "token失效了", nil))
		return
	}
	c.Next()
}
