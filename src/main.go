package main

import (
	"github.com/evndroo/src/controllers/books"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	books.Configure(server)

	server.Run()
}
