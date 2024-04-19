package main

import (
	"github.com/gin-gonic/gin"
	"my-app/src/routes"
	"my-app/src/models"
)

func setupWebApp() *gin.Engine {
	r := gin.Default()
	models.SetupDBConnection();
	routes.SetupControllers(r)
	return r
}

func main() {
	setupWebApp().Run(":80")
}
