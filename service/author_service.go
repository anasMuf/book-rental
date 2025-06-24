package service

import (
	"book-rental-app/model"
	"book-rental-app/repository"
)

type AuthorService interface {
	GetAllAuthor() ([]model.Author, error)
	CreateAuthor(req model.Author) (*model.Author, error)
}

type authorService struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {
	return &authorService{authorRepo}
}

func (s *authorService) GetAllAuthor() ([]model.Author, error) {
	return s.authorRepo.GetAll()
}
func (s *authorService) CreateAuthor(req model.Author) (*model.Author, error) {
	err := s.authorRepo.Create(&req)
	return &req, err
}
