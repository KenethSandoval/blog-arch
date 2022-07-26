package app

import (
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/mongodb"
	echo "github.com/labstack/echo/v4"
)

func New(cfg *config.Config, dbm *mongodb.MongoDB) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Run!")
	})

	return e
}
