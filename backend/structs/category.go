package structs

type Category struct {
	CategoriesID uint64 `gorm:"column:categories_id;primaryKey"`         // ID de la catégorie (type uint64 pour UNSIGNED)
	Name         string `gorm:"column:categories_name"`                  // Nom de la catégorie
	Description  string `gorm:"column:categories_description;type:text"` // Description de la catégorie
}
