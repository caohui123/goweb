package router

import (
	"github.com/gin-gonic/gin"
)

// AdminRoutes 实现 RouteRegistrar 接口的管理员路由
type AdminRoutes struct{}

// RegisterRoutes 注册管理员相关的路由
func (ar *AdminRoutes) RegisterRoutes(r *gin.Engine) {
	adminGroup := r.Group("/api/admin")
	{
		adminGroup.GET("/dashboard", adminDashboard)
	}
}

func init() {
	registrars = append(registrars, &AdminRoutes{})
}

func adminDashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "欢迎来到管理员仪表板"})
}
