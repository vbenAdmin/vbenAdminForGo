package Auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goVben/Dao"
	"goVben/Models"
	"goVben/Redis"
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
	loginForm.Password = "" // 不返回密码
	c.JSON(http.StatusOK, Resful.Result(1000, "", loginForm))
}

type RedisTest struct {
	User1 string `form:"user1"`
	user2 string `form:"user2"`
}

func TestRedis(c *gin.Context) {
	var test RedisTest
	if err := c.ShouldBind(&test); err != nil {
		c.JSON(http.StatusOK, Resful.Result(1004, err.Error(), nil))
		return
	}
	err := Redis.Set("user1", test.User1)
	if err != nil {
		fmt.Println("怎么报错了redis.set", err.Error())
	}
	get, err := Redis.Get("user1")
	if err != nil {
		fmt.Println("怎么报错了redis.get", err.Error())
	}
	fmt.Println("redis.get", get)
	c.JSON(http.StatusOK, Resful.Result(1000, "", test))
}
