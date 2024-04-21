package cmd

import (
	"github.com/gin-gonic/gin"
	"my-app/internal/adapter/input/routes"
	"my-app/internal/adapter/input/models"
)

func SetupWebApp() *gin.Engine {
	r := gin.Default()
	models.SetupDBConnection()
	routes.SetupControllers(r)
	return r
}
