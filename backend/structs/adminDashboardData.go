package structs

import "time"

type AdminDashboardData struct {
	DashboardID   uint64    `gorm:"column:dashboardID;primaryKey;autoIncrement"`
	AdminID       uint64    `gorm:"column:adminID;not null"`
	TotalUsers    uint64    `gorm:"column:totalUsers;default:0"`
	TotalPosts    uint64    `gorm:"column:totalPosts;default:0"`
	TotalComments uint64    `gorm:"column:totalComments;default:0"`
	TotalGuests   uint64    `gorm:"column:totalGuests;default:0"`
	LastLogin     time.Time `gorm:"column:lastLogin"`
	GeneratedAt   time.Time `gorm:"column:generated_at;autoCreateTime"`

}
