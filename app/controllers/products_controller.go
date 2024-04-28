package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/gin-contrib/requestid"
	"my-app/app/models"
	"my-app/app/input_dtos"
	"my-app/app/use_cases"
)

var (
	g = galidator.New().CustomMessages(galidator.Messages{
    "required": "$field is required",
  })
	validator = g.Validator(input_dtos.CreateProductInputDTO{})
)

func ProductsController(r *gin.RouterGroup) {
	indexProducts(r)
	showProduct(r)
	createProduct(r)
	updateProduct(r)
	deleteProduct(r)
}

func indexProducts(r *gin.RouterGroup) {
	var products []models.Product
	models.DB.Find(&products)

	r.GET("/products", func(c *gin.Context) {
		fmt.Printf("%v Processing ProductsController indexProducts\n", requestid.Get(c))
		c.JSON(http.StatusOK, gin.H{"data": products})
	})
}

func showProduct(r *gin.RouterGroup) {
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

func createProduct(r *gin.RouterGroup) {
	r.POST("/products", func(c *gin.Context) {
		dto := input_dtos.CreateProductInputDTO{}

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validator.DecryptErrors(err)})
			return
		}

		uc := use_cases.NewCreateProductUseCase(dto)
		product := uc.Execute()

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}

func updateProduct(r *gin.RouterGroup) {
	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var product models.Product

		if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": fmt.Sprintf("Product with id %v not found", id)}})
			return
		}

		var dto input_dtos.UpdateProductInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validator.DecryptErrors(err)})
			return
		}

		uc := use_cases.NewUpdateProductUseCase(dto, product)
		product = uc.Execute()

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}

func deleteProduct(r *gin.RouterGroup) {
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var product models.Product

		if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"message": fmt.Sprintf("Product with id %v not found", id)}})
			return
		}

		uc := use_cases.NewDeleteProductUseCase(product)
		product = uc.Execute()

		c.JSON(http.StatusOK, gin.H{"data": product})
	})
}
