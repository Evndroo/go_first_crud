package router

import (
	"context"

	"github.com/evndroo/cmd/books/use_cases"
	"github.com/gin-gonic/gin"
)

func Configure(server *gin.Engine, ctx context.Context) {
	booksServer := server.Group("/books")

	booksServer.GET("/", use_cases.WithContextGetAllBooks(ctx))
	booksServer.GET("/:id", use_cases.WithContextGetBooksById(ctx))
	booksServer.POST("/", use_cases.WithContextCreateBook(ctx))
	booksServer.PUT("/:id", use_cases.WithContextUpdateBook(ctx))
	booksServer.DELETE("/:id", use_cases.WithContextDeleteBook(ctx))

}
