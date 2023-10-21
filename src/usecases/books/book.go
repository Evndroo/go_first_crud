package books

import (
	"context"
	"log"
	"net/http"

	"github.com/evndroo/src/entities"
	"github.com/evndroo/src/usecases/context/utils"

	"github.com/gin-gonic/gin"
)

func WithContextGetBooksById(ctx context.Context) gin.HandlerFunc {
	db, success := utils.GetDbFromContext(ctx)
	errorMessages, _ := utils.GetErrorMessagesFromContext(ctx)

	return func(c *gin.Context) {

		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		var book entities.Books
		db.Find(&book, c.Param("id"))

		if book.ID == 0 {
			c.Status(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":     book.ID,
			"title":  book.Title,
			"author": book.Author,
		})

	}
}

func WithContextGetAllBooks(ctx context.Context) gin.HandlerFunc {
	db, success := utils.GetDbFromContext(ctx)
	errorMessages, _ := utils.GetErrorMessagesFromContext(ctx)

	return func(c *gin.Context) {

		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		var books []entities.Books
		result := db.Find(&books)

		if result.Error != nil {
			log.Fatalln("Error getting books: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		response := []gin.H{}

		for _, book := range books {
			response = append(response, gin.H{
				"id":     book.ID,
				"title":  book.Title,
				"author": book.Author,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}

type BookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func WithContextCreateBook(ctx context.Context) gin.HandlerFunc {
	db, success := utils.GetDbFromContext(ctx)
	errorMessages, _ := utils.GetErrorMessagesFromContext(ctx)

	return func(c *gin.Context) {
		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		var body BookDTO

		err := c.ShouldBindJSON(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		book := &entities.Books{
			Title:  body.Title,
			Author: body.Author,
		}

		result := db.Create(book)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":     book.ID,
			"title":  book.Title,
			"author": book.Author,
		})
	}
}
