package service

import (
	"book-rental-app/model"
	"book-rental-app/repository"
)

type BookService interface {
	GetBook(param int) ([]model.Book, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo}
}

func (s *bookService) GetBook(param int) ([]model.Book, error) {
	return s.bookRepo.Get(param)
}
