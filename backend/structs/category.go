package structs

type Category struct {
	CategoriesID          uint64 `gorm:"column:categories_id;primaryKey;autoIncrement"`
	CategoriesName        string `gorm:"column:categories_name;unique;size:255"`
	CategoriesDescription string `gorm:"column:categories_description;type:text"`
}
