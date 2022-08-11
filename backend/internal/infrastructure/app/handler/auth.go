package handler

import (
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/middle"
	"github.com/KenethSandoval/doc-md/pkg/bcrypt"
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

	user, _ := h.atu.GetUserByUsername(ctx, payload.Username)
	if user != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username already exists")
	}

	passwordEncrypt, err := bcrypt.HashPassword(payload.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	payload.ID = primitive.NewObjectID()
	payload.Password = passwordEncrypt

	if err := h.atu.RegisterUser(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "Register successfully")
}

func (h *Handler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.LoginPayload{}
	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.atu.LoginUser(ctx, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Credencials invalida")
	}

	token, err := middle.CreateJWTToken("secret", user)

	return c.JSON(http.StatusOK, token)
}
