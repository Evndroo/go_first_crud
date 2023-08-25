package books

import (
	"github.com/evndroo/src/usecases/books"

	"github.com/gin-gonic/gin"
)

func Configure(server *gin.Engine) {
	server.GET("/books", books.GetBooks)
}
