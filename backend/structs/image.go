package structs

type Image struct {
	ImageID  uint   `gorm:"column:images_id;primaryKey;autoIncrement"` // Identifiant unique (clé primaire, auto-incrément)
	URL      string `gorm:"column:url"`                                // URL de l'image
	Filename string `gorm:"column:filename"`                           // Nom du fichier
	Data     []byte `gorm:"column:data"`                               // Données de l'image (BLOB)
}
