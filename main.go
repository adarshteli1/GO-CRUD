package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    int    `json:"id" `
	Name  string `json:"name" `
	Type  string `json:"type" `
	Price int    `json:"price"`
}

var books = []book{
	{ID: 1, Name: "Goblins", Type: "Fiction", Price: 199},
	{ID: 2, Name: "BlackHoles", Type: "Sci-fic", Price: 7000},
	{ID: 3, Name: "Faults", Type: "Drama", Price: 101},
}

func createBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func deleteBook(c *gin.Context) {

	id := c.Param("id")

	for index, book := range books {

		if fmt.Sprint(book.ID) == id {

			books = append(books[:index], books[index+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "Book deleted",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
	})
}

func main() {
	router := gin.Default()
	router.GET("/books", getbooks)
	router.DELETE("/books/:id", deleteBook)
	router.POST("/books", createBooks)
	router.Run(("localhost:8080"))
}
