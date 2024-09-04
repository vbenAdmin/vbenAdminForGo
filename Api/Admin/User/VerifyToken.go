package User

import (
	"github.com/gin-gonic/gin"
	"goVben/Utils/Resful"
	"net/http"
)

func VerifyToken(c *gin.Context) {
	c.JSON(http.StatusOK, Resful.Result(1000, "", nil))
}
