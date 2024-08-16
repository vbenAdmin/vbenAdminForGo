package AdminRouters

import (
	"github.com/gin-gonic/gin"
	"goVben/Api/Admin/Auth"
	"goVben/Controllers"
)

func SetupAdminGroupRouter(r *gin.RouterGroup) {
	adminGroup := r.Group("admin", Controllers.AdminMiddleware)
	setupAdminAuthGroupRouter(adminGroup)
}
func setupAdminAuthGroupRouter(r *gin.RouterGroup) {
	authGroup := r.Group("auth", Controllers.AdminMiddleware)
	authGroup.POST("login", Auth.LoginAPI) // 微信登录接口
}
