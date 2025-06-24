package model

type Genre struct {
	GenreID uint   `gorm:"primaryKey;autoIncrement;column:genre_id" json:"genre_id"`
	Name    string `json:"name"`

	// Books []Book `gorm:"foreignKey:GenreID" json:"books"`
	Books []Book `json:"books"`

	Model
}

func (Genre) TableName() string {
	return "genres"
}
