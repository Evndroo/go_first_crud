package books

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evndroo/cmd/books"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupSuite() {
	// gin.EnvGinMode = gin.TestMode

}

func TestGetAllBooks(t *testing.T) {
	server := gin.Default()
	ctx := context.Background()

	books.Configure(server, ctx)
	httptest.NewRecorder()

	fmt.Println(server)
	fmt.Println("chegou aqui")

	response, err := http.Get("/books")

	fmt.Println(response)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode, "OK response is expected")
}
