package main

import (
	"github.com/caohui123/goweb/internal/router"
)

func main() {
	apiRouter := router.NewAPIRouter()
	apiRouter.Run(":8080")
}
