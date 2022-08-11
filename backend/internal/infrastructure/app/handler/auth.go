package handler

import (
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.RegisterPayload{}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	payload.ID = primitive.NewObjectID()

	if err := h.atu.RegisterUser(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "Register successfully")
}
