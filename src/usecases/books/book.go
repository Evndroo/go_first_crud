package books

import (
	"context"
	"log"
	"net/http"

	"github.com/evndroo/src/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WithContextGetBooks(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, success := ctx.Value("db").(*gorm.DB)

		if !success {
			log.Fatalln("Error getting db from context")
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Sorry, we have a problem, please try again later.",
			})
			return
		}

		var books entities.BookModel
		result := db.Find(&books)

		if result.Error != nil {
			log.Fatalln("Error getting books: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Sorry, we have a problem, please try again later.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":     books.ID,
			"title":  books.Title,
			"author": books.Author,
		})
	}
}
