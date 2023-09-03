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
	return func(c *gin.Context) {
		db, success := utils.GetDbFromContext(ctx)

		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Sorry, we have a problem, please try again later.",
			})
			return
		}

		var book entities.Books
		db.Find(&book, c.Param("id"))

		if book.ID == 0 {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "No book where found",
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

func WithContextGetBooks(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, success := utils.GetDbFromContext(ctx)

		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Sorry, we have a problem, please try again later.",
			})
			return
		}

		var books []entities.Books
		result := db.Find(&books)

		if result.Error != nil {
			log.Fatalln("Error getting books: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Sorry, we have a problem, please try again later.",
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
