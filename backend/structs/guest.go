package structs

import (
	"time"
)

// vérifier si les info dse guest passent bien à user quand on se connecte ou se crée un compte
type Guest struct {
	GuestID            uint64
	GuestCreatedAt     time.Time
	GuestLastVisitedAt *time.Time
}
