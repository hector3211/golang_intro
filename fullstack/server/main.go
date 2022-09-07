package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/controllers"
	"server/initial"
)

func init() {
	initial.ConnectDB()
	initial.LoadEnvs()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todo/first", controllers.GetAllTodos)
	r.GET("/allbooks", controllers.GetAllBooks)
	r.GET("/book/:title", controllers.BookByTitile)

	r.PUT("/book/update/:id", controllers.BookUpdate)
	r.PUT("/price/update/:id", controllers.BookUpdatePrice)

	r.DELETE("/book/delete/:id", controllers.DeleteBookById)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
