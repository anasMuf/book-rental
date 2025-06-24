package handler

import (
	"book-rental-app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(as service.BookService) *BookHandler {
	return &BookHandler{as}
}

func (h *BookHandler) GetBook(c echo.Context) error {
	paramGenreStr := c.QueryParam("genre")
	paramGenre, err := strconv.Atoi(paramGenreStr)
	if err != nil {
		paramGenre = 0
	}
	books, err := h.BookService.GetBook(paramGenre)
	if err != nil || len(books) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available book"})
	}
	return c.JSON(http.StatusOK, books)
}
