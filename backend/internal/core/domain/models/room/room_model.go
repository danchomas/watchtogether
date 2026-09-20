package room

import (
	"time"
)

type Room struct {
	Id           string
	Title        string
	Code         string
	Visibility   Visibility
	PasswordHash string
	CreatedAt    time.Time
}
