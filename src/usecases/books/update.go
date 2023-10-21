package books

import (
	"context"
	"net/http"
	"strconv"

	"github.com/evndroo/src/entities"
	"github.com/evndroo/src/usecases/context/utils"
	"github.com/gin-gonic/gin"
)

type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func WithContextUpdateBook(ctx context.Context) gin.HandlerFunc {
	db, success := utils.GetDbFromContext(ctx)
	errorMessages, _ := utils.GetErrorMessagesFromContext(ctx)

	return func(c *gin.Context) {
		if !success {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errorMessages.InternalServerError,
			})
			return
		}

		var body UpdateBookDTO

		bodyValidationErr := c.ShouldBindJSON(&body)

		if bodyValidationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": bodyValidationErr.Error(),
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

		book := &entities.Books{}

		db.First(&book, id)

		if book.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No book found with this id.",
			})
			return
		}

		result := db.Model(&book).Updates(body)

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
