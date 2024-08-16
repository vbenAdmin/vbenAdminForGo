package Routers

import (
	"github.com/gin-gonic/gin"
	"goVben/Controllers"
	"goVben/Routers/AdminRouters"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized := r.Group("/", Controllers.BasicAuthMiddleware)
	//authorized.POST("admin", func(c *gin.Context) {
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//	if c.Bind(&json) == nil {
	//		c.JSON(http.StatusOK, gin.H{"status": json})
	//	}
	//})
	AdminRouters.SetupAdminGroupRouter(authorized)
	return r
}
