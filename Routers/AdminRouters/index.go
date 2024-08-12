package AdminRouters

import (
	"github.com/gin-gonic/gin"
	"goVben/Controllers"
)

func SetupAdminGroupRouter(r *gin.Engine) {
	adminGroup := r.Group("/", Controllers.AdminMiddleware)
	adminGroup.POST("login", Login.WechatLogin) // 微信登录接口
}
