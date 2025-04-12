package structs
//donc on prendra une image avec les même fonctions que pour les postes mais on les réadaptera facilement
type Image struct {
	ImageID  uint64 `gorm:"column:imageID;primaryKey;autoIncrement"`
	URL      string `gorm:"column:url;size:255"`
	Filename string `gorm:"column:filename;size:255"`
	Data     []byte `gorm:"column:data"`
}
