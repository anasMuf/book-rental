package repository

import (
	"book-rental-app/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Get(param int) ([]model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Get(param int) ([]model.Book, error) {
	var books []model.Book
	query := r.db

	if param != 0 {
		query = query.Where("genre_id = ?", param)
	}

	err := query.
		Preload("Genre").
		Preload("Author").
		Find(&books).Error
	return books, err
}
