# Go gin api endpoint project

## In this project i use the gin framework to grab Pokemon from the pokemon api

## Starting code in main.go

```go:
  import (
       "github.com/gin-gonic/gin" // gin framework
       "net/http"                 // gin dependency
       "requests/requests"        // folder thats holds our routes
     )

     func main() {
      // se up our paths
      r := gin.Default()
      r.SetTrustedProxies([]string{"127.0.0.1"})
      // example path
      r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "pong",
        })
      })
      // paths
      r.GET("/hello", requests.FirstRequest)
      r.GET("/pokemon", requests.SecondRequest)
      // run our server
      r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
     }
```
