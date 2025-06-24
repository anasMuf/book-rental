package repository

import (
	"book-rental-app/model"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAll() ([]model.Genre, error)
	Create(req *model.Genre) error
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{db}
}

func (r *genreRepository) GetAll() ([]model.Genre, error) {
	var genres []model.Genre
	err := r.db.Find(&genres).Error
	return genres, err
}
func (r *genreRepository) Create(req *model.Genre) error {
	return r.db.Create(req).Error
}
