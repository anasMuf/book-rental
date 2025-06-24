package handler

import (
	"book-rental-app/model"
	"book-rental-app/service"
	"net/http"

	"github.com/labstack/echo"
)

type AuthorHandler struct {
	AuthorService service.AuthorService
}

func NewAuthorHandler(as service.AuthorService) *AuthorHandler {
	return &AuthorHandler{as}
}

func (h *AuthorHandler) GetAllAuthor(c echo.Context) error {
	authors, err := h.AuthorService.GetAllAuthor()
	if err != nil || len(authors) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available author"})
	}
	return c.JSON(http.StatusOK, authors)
}

func (h *AuthorHandler) CreateAuthor(c echo.Context) error {
	var req model.Author
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json"})
	}

	author, err := h.AuthorService.CreateAuthor(req)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, author)
}
