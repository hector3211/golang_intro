package controllers

import (
	"github.com/gin-gonic/gin"
	"go_postgres/initial"
	"go_postgres/models"
)

func ProductsCreate(c *gin.Context) {
	var body struct {
		Name     string
		Quantity int
	}
	c.Bind(&body)
	product := models.Product{Name: body.Name, Quantity: body.Quantity}

	result := initial.DB.Create(&product)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}

func AllProducts(c *gin.Context) {
	var products []models.Product
	initial.DB.Find(&products)
	c.JSON(200, gin.H{
		"All Products": products,
	})
}

func ProductByIndex(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	initial.DB.First(&product, id)
	c.JSON(200, gin.H{
		"Product by index": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var body struct {
		Quantity int
	}
	c.Bind(&body)
	id := c.Param("id")
	var product models.Product
	initial.DB.First(&product, id)
	initial.DB.Model(&product).Update(
		"Quantity", body.Quantity)
	c.JSON(200, gin.H{
		"updated": product,
	})
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	initial.DB.Delete(&models.Product{}, id)
	c.JSON(200, gin.H{
		"Deleted": "deleted product of id:" + id,
	})
}
