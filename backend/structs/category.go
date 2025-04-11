package structs

type Category struct {
	CategoriesID   uint64 `gorm:"column:categories_id;primaryKey;autoIncrement"`
	Name           string `gorm:"column:categories_name;size:255"`
	Description    string `gorm:"column:categories_description;type:text"`
}

