package controllers

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func ZipcodesController(r *gin.RouterGroup) {
	indexZipcodes(r)
}

func indexZipcodes(r *gin.RouterGroup) {
	r.GET("/zipcodes", func(c *gin.Context) {
		fmt.Printf("%v Processing ZipcodesController indexZipcodes\n", requestid.Get(c))

		zipcode := c.Query("zipcode")

		if len(zipcode) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"zipcode": "zipcode is required as query string param"}})
			return
		}

		body, response, err := searchZipcode(zipcode)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return
		}

		c.JSON(response.StatusCode(), body)
	})
}

func searchZipcode(zipcode string) (interface{}, *resty.Response, error) {
	var parsedResponseBody interface{}
	response, err := resty.New().R().
		SetResult(&parsedResponseBody).
		Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json", zipcode))

	return parsedResponseBody, response, err
}
