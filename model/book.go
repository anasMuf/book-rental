package model

type Book struct {
	BookID  uint `gorm:"primaryKey;autoIncrement;column:book_id" json:"book_id"`
	GenreID uint `gorm:"column:genre_id" json:"genre_id"`
	// Genre             Genre  `gorm:"foreignKey:GenreID" json:"genre"`
	Genre    Genre `json:"genre"`
	AuthorID uint  `gorm:"column:author_id" json:"author_id"`
	// Author            Author `gorm:"foreignKey:AuthorID" json:"author"`
	Author            Author `json:"author"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	MinAgeRestriction int    `json:"min_age_restriction"`
	CoverUrl          string `json:"cover_url"`

	// Loans []Loan `gorm:"foreignKey:BookID" json:"loans"`
	Loans []Loan `json:"loans"`

	Model
}

func (Book) TableName() string {
	return "books"
}
