package repository

import (
	"book-rental-app/model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAll() ([]model.Author, error)
	Create(req *model.Author) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) GetAll() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Find(&authors).Error
	return authors, err
}
func (r *authorRepository) Create(req *model.Author) error {
	return r.db.Create(req).Error
}
