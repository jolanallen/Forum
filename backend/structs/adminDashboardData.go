package structs

import "time"

// AdminDashboardData represents the data related to the admin dashboard.
// It contains various statistics and information about the admin's account and activity.
type AdminDashboardData struct {
	// DashboardID is the unique identifier for the dashboard data.
	DashboardID uint64
	
	// AdminID is the unique identifier for the admin to whom this dashboard data belongs.
	AdminID uint64
	
	// TotalUsers is the total number of registered users in the system.
	TotalUsers uint64
	
	// TotalPosts is the total number of posts made in the system.
	TotalPosts uint64
	
	// TotalComments is the total number of comments made in the system.
	TotalComments uint64
	
	// TotalGuests is the total number of guests visiting the platform.
	TotalGuests uint64
	
	// LastLogin is the timestamp of the admin's last login to the system.
	LastLogin time.Time
	
	// GeneratedAt is the timestamp when this dashboard data was generated.
	GeneratedAt time.Time
}
