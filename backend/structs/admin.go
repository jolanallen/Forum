package structs

type Admin struct {
	AdminID           uint64
	AdminUsername     string
	AdminPasswordHash string
	AdminEmail        string
}
