package structs

// Image represents an image file stored in the system.
// It contains details about the image, including its unique ID, URL, filename, and the binary data of the image.
type Image struct {
	// ImageID is the unique identifier for the image.
	ImageID uint64
	
	// URL is the accessible path where the image can be retrieved.
	URL string
	
	// Filename is the name of the image file.
	Filename string
	
	// Data contains the raw binary data of the image.
	Data []byte
}
