package models

import "time"

type User struct {
	ID        uint64
	Username  string
	Email     string
	CreatedAt time.Time
}
