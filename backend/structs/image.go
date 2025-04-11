package structs

type Image struct {
	ImageID  uint   `gorm:"column:images_id;primaryKey;autoIncrement"`
	URL      string `gorm:"column:url;size:255"`
	Filename string `gorm:"column:filename;size:255"`
	Data     []byte `gorm:"column:data"`
}
