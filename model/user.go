package model

import (
	"time"
)

type User struct {
	UserID        uint       `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `gorm:"unique;not null" json:"email"`
	Address       string     `json:"address"`
	DateOfBirth   string     `gorm:"type:date" json:"date_of_birth"`
	Password      string     `gorm:"not null" json:"password"`
	LastLoginDate *time.Time `json:"last_login_date"`
	JwtToken      string     `json:"jwt_token"`
	// Loans         []Loan    `gorm:"foreignKey:UserID" json:"loans"`
	Loans []Loan `json:"loans"`
	Model
}

func (User) TableName() string {
	return "users"
}
