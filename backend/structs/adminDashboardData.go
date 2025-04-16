package structs

import "time"

type AdminDashboardData struct {
	DashboardID   uint64
	AdminID       uint64
	TotalUsers    uint64
	TotalPosts    uint64
	TotalComments uint64
	TotalGuests   uint64
	LastLogin     time.Time
	GeneratedAt   time.Time
}
