package Auth

import (
	"github.com/gin-gonic/gin"
	"goVben/Dao"
	"goVben/Models"
	"goVben/Redis"
	"goVben/Utils/Config"
	"goVben/Utils/Jwt"
	"goVben/Utils/Resful"
	"net/http"
)

func LoginAPI(c *gin.Context) {
	var loginForm Models.Users
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusOK, Resful.Result(1004, err.Error(), nil))
		return
	}
	err := Dao.GetLoginUser(&loginForm)
	if err != nil {
		c.JSON(http.StatusOK, Resful.Result(1005, err.Error(), nil))
		return
	}
	token, expAt, err := Jwt.GenerateJWT(loginForm.ID, Config.Conf.JwtKey)
	if err != nil {
		c.JSON(http.StatusOK, Resful.Result(1005, err.Error(), nil))
		return
	}
	loginForm.Token = token
	Redis.SetUserToken(loginForm.ID, expAt)
	loginForm.Password = "" // 不返回密码
	c.JSON(http.StatusOK, Resful.Result(1000, "", loginForm))
}

type RedisTest struct {
	User1 string `form:"user1"`
	user2 string `form:"user2"`
}
