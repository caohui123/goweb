package router

import "github.com/gin-gonic/gin"

type APIRouter struct {
	Router *gin.Engine
}

func NewAPIRouter(r *gin.Engine) *gin.Engine {
	api := &APIRouter{
		Router: r,
	}
	return api.initRoutes(r)
}

func (api *APIRouter) initRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", api.ping)
	RegisterRoutes(r)
	return api.Router
}

func (api *APIRouter) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
