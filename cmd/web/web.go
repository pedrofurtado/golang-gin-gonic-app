package web

import (
	"github.com/gin-gonic/gin"
	"my-app/app/routes"
)

func SetupWebApp() *gin.Engine {
	r := gin.Default()
	routes.SetupControllers(r)
	return r
}
