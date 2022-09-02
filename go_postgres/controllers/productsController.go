package controllers

import (
	"github.com/gin-gonic/gin"
	"go_postgres/initial"
	"go_postgres/models"
)

/*
this function takes in a body (name and quantity)
like our model
*/
func ProductsCreate(c *gin.Context) {
	var body struct {
		Name     string
		Quantity int
	}
	c.Bind(&body) // this binds our url path to the body
	/*
	   we create a product variable that is an instance of our
	   Porduct model and we "bind" our body params to it
	*/
	product := models.Product{Name: body.Name, Quantity: body.Quantity}
	// we create the product we .Create
	result := initial.DB.Create(&product)
	// result returns an error if there is one
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}

// This function returns all products in our db
func AllProducts(c *gin.Context) {
	// we make a products varaible thats a list of our models
	var products []models.Product
	// we call DB.Find to grab all
	initial.DB.Find(&products)
	c.JSON(200, gin.H{
		"All Products": products,
	})
}

// This function retruns a product by index specified
func ProductByIndex(c *gin.Context) {
	// we grab the id param from the path
	id := c.Param("id")
	// we grab our Product table/model
	var product models.Product
	// we query for the id
	initial.DB.First(&product, id)
	c.JSON(200, gin.H{
		"Product by index": product,
	})
}

// This function updates a product
func UpdateProduct(c *gin.Context) {
	// Lets update the quantity
	var body struct {
		Quantity int
	}
	c.Bind(&body)
	// grab id param from path
	id := c.Param("id")
	var product models.Product
	// we query for the selected product
	initial.DB.First(&product, id)
	// we updated the quantity column of selected product
	initial.DB.Model(&product).Update(
		"Quantity", body.Quantity)
	c.JSON(200, gin.H{
		"updated": product,
	})
}

// This function deletes a selected product by id
func DeleteProductById(c *gin.Context) {
	id := c.Param("id") // we grab the id from the path/url
	// here we call .Delete pass or Product model object/table
	// and id
	initial.DB.Delete(&models.Product{}, id)
	c.JSON(200, gin.H{
		"Deleted": "deleted product of id:" + id,
	})
}
