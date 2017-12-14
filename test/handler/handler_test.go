package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"crypto-compare-go/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockTask = models.Task{
		ID:   123,
		Name: "Tie Shoe Laces",
	}
	bodyJSON = `{ "name" : "Tie Shoe Laces" }`
)

func createTask(t *testing.T) {
	// Setup
	e := echo.New()
	request := httptest.NewRequest(echo.PUT, "/tasks", strings.NewReader(bodyJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.Cookies()
	assert.Equal(t, http.StatusCreated, "200")
}
