package AdminRouters

import (
	"github.com/gin-gonic/gin"
	"goVben/Api/Admin/Auth"
	"goVben/Api/Admin/User"
	"goVben/Controllers"
)

func SetupAdminGroupRouter(r *gin.RouterGroup) {
	adminGroup := r.Group("admin", Controllers.AdminMiddleware)
	adminGroup.POST("login", Auth.LoginAPI) // 微信登录接口
	setupAdminAuthGroupRouter(adminGroup)
}
func setupAdminAuthGroupRouter(r *gin.RouterGroup) {
	userGroup := r.Group("user", Controllers.VerifyTokenController)
	userGroup.POST("verifyToken", User.VerifyToken)
}
