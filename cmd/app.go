package main

import (
	"github.com/caohui123/goweb/internal/router"
	"github.com/caohui123/goweb/pkg/config"
	"github.com/caohui123/goweb/pkg/logger"
	"github.com/caohui123/goweb/pkg/server"
	vaildator_zh "github.com/caohui123/goweb/pkg/validator"
	"github.com/gin-gonic/gin"
	"log"
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
	r = router.NewAPIRouter(r)
	server.Run(r, conf)
}
