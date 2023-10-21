package books

import (
	"context"

	"github.com/evndroo/src/usecases/books"

	"github.com/gin-gonic/gin"
)

func Configure(server *gin.Engine, ctx context.Context) {
	booksServer := server.Group("/books")

	booksServer.GET("/", books.WithContextGetAllBooks(ctx))
	booksServer.GET("/:id", books.WithContextGetBooksById(ctx))
	booksServer.POST("/", books.WithContextCreateBook(ctx))
	booksServer.PUT("/:id", books.WithContextUpdateBook(ctx))
}
