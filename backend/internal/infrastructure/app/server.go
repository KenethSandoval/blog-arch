package app

import (
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/config"
	echo "github.com/labstack/echo/v4"
)

func New(cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Run!")
	})

	return e
}
