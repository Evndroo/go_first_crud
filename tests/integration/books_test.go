package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evndroo/cmd/books/config"
	"github.com/evndroo/cmd/books/router"
	"github.com/evndroo/cmd/context/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupSuite(t *testing.T) {
	httptest.NewRecorder()
	ctx := context.Background()
	server := gin.Default()
	db, err := config.ConnectDB()

	if err != nil {
		t.FailNow()
	}

	ctx = utils.WithDbContext(ctx, db)
	ctx = utils.WithErrorMessagesContext(ctx)

	router.Configure(server, ctx)
}

func TestGetAllBooks(t *testing.T) {
	response, err := http.Get("/books")

	fmt.Println(response)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode, "OK response is expected")
}
