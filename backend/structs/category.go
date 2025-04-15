package structs

type Category struct {
	CategoryID          uint64 `gorm:"column:categoryID;primaryKey;autoIncrement"`
	CategoryName        string `gorm:"column:categoryName;unique;size:255"`
	CategoryDescription string `gorm:"column:categoryDescription;type:text"`
}
