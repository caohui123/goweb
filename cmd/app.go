package main

import (
	"github.com/caohui123/goweb/router"
)

func main() {
	apiRouter := router.NewApiRouter()
	apiRouter.Run(":8080")
}
