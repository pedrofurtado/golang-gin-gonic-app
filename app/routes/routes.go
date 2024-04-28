package routes

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/requestid"
	"github.com/google/uuid"
	"my-app/app/controllers"
)

func SetupControllers(r *gin.Engine) {
	setupRootRoute(r)
	setupApiRoutes(r)
}

func setupRootRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "My App")
	})
}

func setupApiRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	setupApiMiddlewares(apiGroup)
	controllers.ProductsController(apiGroup)
	controllers.ZipcodesController(apiGroup)
}

func setupApiMiddlewares(r *gin.RouterGroup) {
	setupRequestIDMiddleware(r)
	setupBasicAuthMiddleware(r)
	setupRequestAndResponseLogMiddleware(r)
}

func setupRequestIDMiddleware(r *gin.RouterGroup) {
	r.Use(
    requestid.New(
      requestid.WithGenerator(func() string {
        return uuid.New().String()
      }),
      requestid.WithCustomHeaderStrKey("X-Request-Id"),
    ),
  )
}

func setupBasicAuthMiddleware(r *gin.RouterGroup) {
	// Basic auth
	// To authenticate, set header "Authorization" with value "Basic YWRtaW46MTIzNA=="

	allowedAccounts := gin.Accounts{
		"admin": "1234",
	}

	r.Use(gin.BasicAuth(allowedAccounts))
}

func setupRequestAndResponseLogMiddleware(r *gin.RouterGroup) {
	// Request logging
	r.Use(func(c *gin.Context) {
		fmt.Printf("%s Starting request %s %s %s %s %s at %s\n",
			requestid.Get(c),
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			c.Request.Host,
			c.Request.RemoteAddr,
			time.Now(),
		)
		c.Next()
	})

	// Response logging
	r.Use(func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Printf("%s Finishing request %s %s %s %s %s in %s at %s\n",
			requestid.Get(c),
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			c.Request.Host,
			c.Request.RemoteAddr,
			latency,
			time.Now(),
		)
	})
}
