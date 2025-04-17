package structs

// Admin represents an administrator in the system.
// It contains the ID, username, password hash, and email of the admin.
type Admin struct {
	// AdminID is the unique identifier for the admin.
	AdminID uint64
	
	// AdminUsername is the username of the admin.
	AdminUsername string
	
	// AdminPasswordHash is the hashed password of the admin.
	AdminPasswordHash string
	
	// AdminEmail is the email address of the admin.
	AdminEmail string
}
