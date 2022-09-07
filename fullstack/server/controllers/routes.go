package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/initial"
	"server/models"
)

func init() {
	initial.GetFirstTodos()
}

func GetAllTodos(c *gin.Context) {
	var data = initial.GetFirstTodos()
	c.JSON(200, gin.H{
		"todos": data,
	})
}

func BooksCreate(c *gin.Context) {
	var body struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
	c.Bind(&body)
	book := models.Book{Title: body.Title, Description: body.Description, Price: body.Price}
	result := initial.DB.Create(&book)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"created": book,
	})
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")
	initial.DB.Delete(&models.Book{}, id)
	c.JSON(http.StatusOK, gin.H{
		"Deleted": "deleted book of id:" + id,
	})
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	initial.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func BookByTitile(c *gin.Context) {
	title := c.Param("title")
	var book models.Book
	initial.DB.First(&book, title)
	c.JSON(http.StatusOK, gin.H{
		"product by title": title,
	})
}

func BookUpdatePrice(c *gin.Context) {
	var body struct {
		Price float64 `json:"price"`
	}
	c.Bind(&body)
	id := c.Param("id")
	var book models.Book
	initial.DB.First(&book, id)
	initial.DB.Model(&book).Update(
		"Price", body.Price)
	c.JSON(http.StatusOK, gin.H{
		"updated": book,
	})
}

func BookUpdate(c *gin.Context) {
	var body struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
	c.Bind(&body)
	id := c.Param("id")
	var book models.Book
	initial.DB.First(&book, id)
	initial.DB.Model(&book).Select(id).Updates(map[string]interface{}{
		"Title":       body.Title,
		"Description": body.Description,
		"Price":       body.Price,
	})
}
