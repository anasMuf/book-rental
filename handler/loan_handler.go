package handler

import (
	"book-rental-app/model"
	"book-rental-app/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type LoanHandler struct {
	LoanService service.LoanService
}

func NewLoanHandler(as service.LoanService) *LoanHandler {
	return &LoanHandler{as}
}

func (h *LoanHandler) GetAllLoan(c echo.Context) error {
	loans, err := h.LoanService.GetAllLoan()
	if err != nil || len(loans) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No available author"})
	}
	return c.JSON(http.StatusOK, loans)
}

func (h *LoanHandler) CreateLoan(c echo.Context) error {
	var req struct {
		BookID       int `json:"book_id"`
		LoanDuration int `json:"loan_duration"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json"})
	}

	userJWT := c.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	now := time.Now()
	tanggalString := now.Format("2006-01-02")
	tanggalBaru := now.AddDate(0, 0, req.LoanDuration).Format("2006-01-02")

	loanReq := model.Loan{
		UserID:     userID,
		BookID:     uint(req.BookID),
		LoanDate:   tanggalString,
		DueDate:    tanggalBaru,
		ReturnDate: nil,
	}

	loan, err := h.LoanService.CreateLoan(loanReq)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, loan)
}
