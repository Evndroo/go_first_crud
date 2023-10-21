package books

import (
	"context"
	"net/http"

	"github.com/evndroo/src/entities"
	"github.com/evndroo/src/use-cases/context/utils"
	"github.com/gin-gonic/gin"
)

func WithContextDeleteBook(context context.Context) gin.HandlerFunc {
	db, success := utils.GetDbFromContext(context)
	errorMessages, _ := utils.GetErrorMessagesFromContext(context)

	return func(c *gin.Context) {
		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessages.InvalidIdBadRequest,
			})
			return
		}

		book := &entities.Books{}

		db.First(&book, id)

		if book.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No book found where found with this id.",
			})
			return
		}

		result := db.Delete(&id)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book deleted successfully.",
		})
	}
}
