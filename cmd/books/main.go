package books

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/evndroo/cmd/books/config"
	"github.com/evndroo/cmd/context/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if os.Getenv("ENVIRONMENT") == "" {
		fmt.Println("ENVIRONMENT is not set, setting envs from .env file")
		godotenv.Load()
	}

	db, connectionError := config.ConnectDB()

	if connectionError != nil {
		log.Fatalln("Error connecting to database: ", connectionError)
		panic(connectionError)
	}

	ctx = utils.WithDbContext(ctx, db)
	ctx = utils.WithErrorMessagesContext(ctx)

	server := gin.Default()

	Configure(server, ctx)

	server.Run()
}
