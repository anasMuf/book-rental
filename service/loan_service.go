package service

import (
	"book-rental-app/model"
	"book-rental-app/repository"
)

type LoanService interface {
	GetAllLoan() ([]model.Loan, error)
	CreateLoan(req model.Loan) (*model.Loan, error)
}

type loanService struct {
	loanRepo repository.LoanRepository
}

func NewLoanService(loanRepo repository.LoanRepository) LoanService {
	return &loanService{loanRepo}
}

func (s *loanService) GetAllLoan() ([]model.Loan, error) {
	return s.loanRepo.GetAll()
}
func (s *loanService) CreateLoan(req model.Loan) (*model.Loan, error) {
	err := s.loanRepo.Create(&req)
	if err != nil {
		return nil, err
	}

	// Ambil ulang dengan relasi User dan Book
	return s.loanRepo.FindByID(req.LoanID)
}
