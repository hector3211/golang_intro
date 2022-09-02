package main

import (
	"github.com/gin-gonic/gin"
	"go_postgres/controllers"
	"go_postgres/initial"
)

func init() {
	initial.LoadEnvs()
	initial.ConnectDB()
}

func main() {
	// we initialize our gin serve
	r := gin.Default()
	// All the routes for our server
	r.GET("/allproducts", controllers.AllProducts)
	r.GET("/products/:id", controllers.ProductByIndex)
	r.POST("/products", controllers.ProductsCreate)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProductById)
	// we tell gin server to run
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
