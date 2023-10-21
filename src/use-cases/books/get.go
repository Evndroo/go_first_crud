package books

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/evndroo/src/entities"
	"github.com/evndroo/src/use-cases/context/utils"
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

		id, idErr := strconv.Atoi(c.Param("id"))

		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessages.InvalidIdBadRequest,
			})
			return
		}

		var book entities.Books
		db.First(&book, id)

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
