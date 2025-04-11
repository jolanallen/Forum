package structs

type Category struct {
	CategoriesID          uint64 `gorm:"column:categoriesID;primaryKey;autoIncrement"`
	CategoriesName        string `gorm:"column:categoriesName;unique;size:255"`
	CategoriesDescription string `gorm:"column:categoriesDescription;type:text"`
}
