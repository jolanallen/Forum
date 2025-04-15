package structs

import (
	"time"
)
//vérifier si les info dse guest passent bien à user quand on se connecte ou se crée un compte
type Guest struct {
	GuestID        uint64           `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time        `gorm:"autoCreateTime"`
	LastVisitedAt  *time.Time

	Sessions       []SessionGuest   `gorm:"foreignKey:GuestID"`
}
