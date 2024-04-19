package routes

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/requestid"
	"github.com/google/uuid"
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
	setupV1Middlewares(v1Group)
	controllersV1.ProductsController(v1Group)
	controllersV1.ZipcodesController(v1Group)
}

func setupV1Middlewares(r *gin.RouterGroup) {
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
