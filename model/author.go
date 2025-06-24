package model

type Author struct {
	AuthorID uint   `gorm:"primaryKey;autoIncrement;column:author_id" json:"author_id"`
	Name     string `json:"name"`

	// Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
	Books []Book `json:"books"`

	Model
}

func (Author) TableName() string {
	return "authors"
}
