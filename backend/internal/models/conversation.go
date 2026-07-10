package models

import "time"

type Conversation struct {
	ID            uint64
	UserIDs       []uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
