package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)
	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Sigit Budi",
		"bio":  "miwon",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Sigit Budi",
		"message": "HALO SEMUA",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BooksInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) {
	//title,price
	var bookInput BooksInput
	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error di kolom %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
