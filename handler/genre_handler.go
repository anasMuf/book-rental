package handler

import (
	"book-rental-app/model"
	"book-rental-app/service"
	"net/http"

	"github.com/labstack/echo"
)

type GenreHandler struct {
	GenreService service.GenreService
}

func NewGenreHandler(as service.GenreService) *GenreHandler {
	return &GenreHandler{as}
}

func (h *GenreHandler) GetAllGenre(c echo.Context) error {
	genres, err := h.GenreService.GetAllGenre()
	if err != nil || len(genres) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available author"})
	}
	return c.JSON(http.StatusOK, genres)
}

func (h *GenreHandler) CreateGenre(c echo.Context) error {
	var req model.Genre
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json"})
	}

	genre, err := h.GenreService.CreateGenre(req)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, genre)
}
