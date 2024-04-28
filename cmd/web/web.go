package web

import (
	"github.com/gin-gonic/gin"
	"my-app/app/routes"
	"my-app/app/models"
)

func SetupWebApp() *gin.Engine {
	r := gin.Default()
	models.SetupDBConnection()
	routes.SetupControllers(r)
	return r
}
