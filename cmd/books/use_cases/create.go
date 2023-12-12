package use_cases

import (
	"context"
	"net/http"

	"github.com/evndroo/cmd/books/entities"
	"github.com/evndroo/cmd/context/utils"

	"github.com/gin-gonic/gin"
)

type CreateBookDTO struct {
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

		var body CreateBookDTO

		err := c.ShouldBindJSON(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		book := entities.Books{
			Title:  body.Title,
			Author: body.Author,
		}

		result := db.Create(&book)

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
