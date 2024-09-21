package router

import "github.com/gin-gonic/gin"

// RouteRegistrar 用于定义注册路由的方法
type RouteRegistrar interface {
	RegisterRoutes(r *gin.Engine)
}

var registrars []RouteRegistrar

// RegisterRoutes 自动注册所有路由
func RegisterRoutes(r *gin.Engine) {

	//registrars := []RouteRegistrar{
	//	&UserRoutes{},
	//	&AdminRoutes{},
	//}

	for _, registrar := range registrars {
		registrar.RegisterRoutes(r)
	}
}
