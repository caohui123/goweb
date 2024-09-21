package router

import (
	"github.com/caohui123/goweb/internal/model"
	vaildator_zh "github.com/caohui123/goweb/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRoutes 实现 RouteRegistrar 接口的用户路由
type UserRoutes struct{}

// RegisterRoutes 注册用户相关的路由
func (ur *UserRoutes) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", getUsers)
		userGroup.POST("/", createUser)
		userGroup.POST("/adduser", AddUser)
	}
}
func init() {
	registrars = append(registrars, &UserRoutes{})
}

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取用户列表"})
}

func createUser(c *gin.Context) {
	c.JSON(201, gin.H{"message": "用户创建成功"})
}
func AddUser(c *gin.Context) {
	var data model.User
	//接收请求参数
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "参数错误",
			"data":    vaildator_zh.ErrorRes(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
	})
	return

}
