package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	controllersV1 "my-app/src/controllers/v1"
)

func SetupControllers(r *gin.Engine) {
	setupRootRoute(r)
	setupV1Routes(r)
}

func setupRootRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "My App")
	})
}

func setupV1Routes(r *gin.Engine) {
	v1Group := r.Group("/v1")
	setupBasicAuthMiddleware(v1Group)
	controllersV1.ProductsController(v1Group)
}

func setupBasicAuthMiddleware(r *gin.RouterGroup) {
	// Basic auth
	// To authenticate, set header "Authorization" with value "Basic YWRtaW46MTIzNA=="

	allowedAccounts := gin.Accounts{
		"admin": "1234",
	}

	r.Use(gin.BasicAuth(allowedAccounts))
}
