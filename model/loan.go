package model

type Loan struct {
	LoanID uint `gorm:"primaryKey;autoIncrement;column:loan_id" json:"loan_id"`
	UserID uint `gorm:"column:user_id" json:"user_id"`
	// User       User           `gorm:"foreignKey:UserID" json:"user"`
	User   User `json:"user"`
	BookID uint `gorm:"column:book_id" json:"book_id"`
	// Book       Book           `gorm:"foreignKey:BookID" json:"book"`
	Book       Book    `json:"book"`
	LoanDate   string  `gorm:"type:date" json:"loan_date"`
	DueDate    string  `gorm:"type:date" json:"due_date"`
	ReturnDate *string `gorm:"type:date" json:"return_date"`

	Model
}

func (Loan) TableName() string {
	return "loans"
}
