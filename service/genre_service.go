package service

import (
	"book-rental-app/model"
	"book-rental-app/repository"
)

type GenreService interface {
	GetAllGenre() ([]model.Genre, error)
	CreateGenre(req model.Genre) (*model.Genre, error)
}

type genreService struct {
	genreRepo repository.GenreRepository
}

func NewGenreService(genreRepo repository.GenreRepository) GenreService {
	return &genreService{genreRepo}
}

func (s *genreService) GetAllGenre() ([]model.Genre, error) {
	return s.genreRepo.GetAll()
}
func (s *genreService) CreateGenre(req model.Genre) (*model.Genre, error) {
	err := s.genreRepo.Create(&req)
	return &req, err
}
