package router

import "github.com/gin-gonic/gin"

type APIRouter struct {
	Router *gin.Engine
}

func NewAPIRouter() *APIRouter {
	router := gin.Default()
	apiRouter := APIRouter{Router: router}
	apiRouter.initRoutes()

	return &apiRouter
}

func (api *APIRouter) initRoutes() {
	api.Router.GET("/ping", api.ping)
}

func (api *APIRouter) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
