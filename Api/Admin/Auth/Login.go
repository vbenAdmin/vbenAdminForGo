package Auth

import (
	"github.com/gin-gonic/gin"
	"goVben/Dao"
	"goVben/Models"
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
	c.JSON(http.StatusOK, Resful.Result(1000, "", loginForm))
}
