package structs



type Comment struct {
	ID          uint      `gorm:"primaryKey"`
	CreatorID   uint      `gorm:"not null"`
	Texte       string    `gorm:"type:text"`
}