package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateArticle(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.ArticleCreatePayload{}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// call usecase of business logic
	payload.ID = primitive.NewObjectID()

	err := createReadmeArticle(payload.Title, payload.Content)
	if err != nil {
		return echo.ErrInternalServerError
	}
	// add timestamp
	if err := h.aru.CreateArticle(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "CREATED")
}

func createReadmeArticle(nameFile, content string) error {
	path := "/home/stivarch/Escritorio/gnu/projects/documentation-md/client/src/_posts/" + nameFile + ".mdx"

	headContent := fmt.Sprintf(`
---
date: '2021-11-25'
title: %s
description: A quick guide into Next.js and Typescript with deployment to vercel
prerequisites: ['Node.js installed on your computer', 'Basic knowledge working with Next.js and TypeScript']
stacks: ['Next.js','TypeScript','Git']
---

<Prerequisites />

<Stacks /> %s`, nameFile, "\n\n")

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	data := []byte(headContent + content)

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
