package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:Password@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database")
	}
	fmt.Println("Koneksi database berhasil")
	//migrasi
	db.AutoMigrate(&book.Book{})

	//CRUD
	book := book.Book{}
	book.Title = "Mein Kampf"
	book.Price = 120000
	book.Discount = 5000
	book.Rating = 4
	book.Decsription = "Kisah Adolf Hitler semasa Hidup"

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("===============")
		fmt.Println("Error buat baru")
		fmt.Println("===============")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	// v2 := router.Group("/v2")

	router.Run()
}
