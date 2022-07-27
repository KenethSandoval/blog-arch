package app_test

import (
	"net/http"
	"testing"

	"github.com/KenethSandoval/doc-md/internal/infrastructure/tester"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorNotFound(t *testing.T) {
	testy := tester.New()

	code, body := testy.RequestWithServe(http.MethodGet, "/xxx", nil, nil)

	assert.Equal(t, http.StatusNotFound, code)
	assert.Equal(t, "{\"message\":\"Not Found\"}\n", body)
}