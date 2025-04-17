package structs

// Category represents a category in the system.
// It holds information related to a specific category used to organize posts.
type Category struct {
	// CategoryID is the unique identifier for the category.
	CategoryID uint64
	
	// CategoryName is the name of the category.
	CategoryName string
	
	// CategoryDescription provides a description of the category.
	CategoryDescription string
}
