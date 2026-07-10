package models

import "time"

type Message struct {
	ID             uint64
	SenderID       uint64
	ReceiverID     uint64
	Content        string
	Encrypted      bool
	CreatedAt      time.Time
}
