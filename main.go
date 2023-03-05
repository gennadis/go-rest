package main

import (
	"errors"
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
	router.GET("/books/:id", bookById)
	router.POST("/", createBook)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
