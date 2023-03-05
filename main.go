package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 1},
	{ID: "2", Title: "Title 2", Author: "Author 2", Quantity: 2},
	{ID: "3", Title: "Title 3", Author: "Author 3", Quantity: 3},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
