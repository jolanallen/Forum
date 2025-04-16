package structs

import "time"

type SessionAdmin struct {
	ASessionID    uint64
	AAdminID      uint64
	ASessionToken string
	AExpiresAt    time.Time
	ACreatedAt    time.Time
}
