package structs

type Admin struct {
	AdminID           uint64 `gorm:"primaryKey;autoIncrement"`
	AdminUsername     string `gorm:"size:255;not null;unique"`
	AdminPasswordHash string `gorm:"size:255;not null"`
	AdminEmail        string `gorm:"size:255;not null;unique"`

	Sessions      []SessionAdmin       `gorm:"foreignKey:AdminID"`
	DashboardData []AdminDashboardData `gorm:"foreignKey:AdminID"`
}
