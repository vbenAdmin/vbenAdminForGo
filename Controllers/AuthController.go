package Controllers

import (
	"github.com/gin-gonic/gin"
	"goVben/Utils/Resful"
	"net/http"
)

func BasicAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusOK, Resful.Result(1001, "", nil))
		return
	}
	c.Next()
}
