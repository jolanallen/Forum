package structs

type Image struct {
    URL      string `gorm:"size:255"`  // Stocke l'URL de l'image
    Filename string `gorm:"size:255"`  // Nom du fichier
    Data     []byte `gorm:"type:blob"` // Stocke l'image en binaire (SQLite/MySQL)
}
