package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/gin-contrib/requestid"
	"my-app/src/models"
)

type CreateProductData struct {
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price"       binding:"required"`
	Quantity    int       `json:"quantity"    binding:"required"`
	Active      bool      `json:"active"`
}

type UpdateProductData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Active      bool      `json:"active"`
}

var (
	g = galidator.New().CustomMessages(galidator.Messages{
    "required": "$field is required",
  })
	validator = g.Validator(CreateProductData{})
)

func ProductsController(r *gin.RouterGroup) {
	index(r)
	show(r)
	create(r)
	update(r)
	delete(r)
}

func index(r *gin.RouterGroup) {
	var products []models.Product
	models.DB.Find(&products)

	r.GET("/products", func(c *gin.Context) {
		fmt.Printf("%v Processing ProductsController index\n", requestid.Get(c))
		c.JSON(http.StatusOK, gin.H{"data": products})
	})
}

func show(r *gin.RouterGroup) {
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var product models.Product

		if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": fmt.Sprintf("Product with id %v not found", id)}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}

func create(r *gin.RouterGroup) {
	r.POST("/products", func(c *gin.Context) {
		params := &CreateProductData{}
		if err := c.BindJSON(params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validator.DecryptErrors(err)})
		} else {
			product := models.Product{
				Name: params.Name,
				Description: params.Description,
				Price: params.Price,
				Quantity: params.Quantity,
				Active: params.Active,
			}
  		models.DB.Create(&product)

			c.JSON(http.StatusOK, gin.H{"data": product})
		}
	})
}

func update(r *gin.RouterGroup) {
	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var product models.Product

		if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": fmt.Sprintf("Product with id %v not found", id)}})
			return
		}

		var input UpdateProductData
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validator.DecryptErrors(err)})
			return
		}

		models.DB.Model(&product).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}

func delete(r *gin.RouterGroup) {
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var product models.Product
		if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": fmt.Sprintf("Product with id %v not found", id)}})
			return
		}

		models.DB.Delete(&product)

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}
