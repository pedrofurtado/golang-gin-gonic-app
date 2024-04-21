package cmd

import (
	"github.com/gin-gonic/gin"
	"my-app/src/routes"
	"my-app/src/models"
)

func SetupWebApp() *gin.Engine {
	r := gin.Default()
	models.SetupDBConnection()
	routes.SetupControllers(r)
	return r
}
