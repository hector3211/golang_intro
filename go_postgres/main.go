package main

import (
	"github.com/gin-gonic/gin"
	"go_postgres/initial"
	"net/http"
)

func init() {
	initial.LoadEnvs()
	initial.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
