package main

import (
	"github.com/caohui123/goweb/pkg/config"
	"github.com/caohui123/goweb/pkg/logger"
	"github.com/caohui123/goweb/pkg/server"
	vaildator_zh "github.com/caohui123/goweb/pkg/validator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	conf := config.Load("")
	err := logger.InitLogger(&conf.LogConfig)
	if err != nil {
		panic(err)
	}
	//mysql.InitMysql(&conf.DBConfig)
	//初始化validator
	//初始化翻译器，这部分可放在main.go或router.go中
	if err := vaildator_zh.InitTrans("zh"); err != nil {
		log.Fatalf("init trans failed, err:%v\n", err)
		return
	}
	r := gin.New()
	r.POST("/adduser", AddUser)
	server.Run(r, conf)
}

type User struct {
	UserName string `json:"username" binding:"required,min=4,max=12" label:"用户名"`
	PassWord string `json:"password" binding:"required,min=6,max=20" label:"密码"`
}

func AddUser(c *gin.Context) {
	var data User
	//接收请求参数
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": vaildator_zh.ErrRespString(err),
			"data":    nil,
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
