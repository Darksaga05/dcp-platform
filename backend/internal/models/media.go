package models

import "time"

type Media struct {
	ID        uint64
	MessageID uint64
	URL       string
	Type      string
	Size      int64
	CreatedAt time.Time
}
