package repository

import (
	"book-rental-app/model"

	"gorm.io/gorm"
)

type LoanRepository interface {
	GetAll() ([]model.Loan, error)
	FindByID(id uint) (*model.Loan, error)
	Create(req *model.Loan) error
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db}
}

func (r *loanRepository) GetAll() ([]model.Loan, error) {
	var loans []model.Loan
	err := r.db.Find(&loans).Error
	return loans, err
}

func (r *loanRepository) FindByID(id uint) (*model.Loan, error) {
	var loan model.Loan
	err := r.db.
		Preload("User").
		Preload("Book").
		Preload("Book.Genre").
		Preload("Book.Author").
		First(&loan, id).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) Create(req *model.Loan) error {
	return r.db.Create(req).Error
}
