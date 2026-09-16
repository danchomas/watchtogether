package room

import "time"

type Room struct {
	Id           string
	Tittle       string
	Code         string
	Visibility   Visibility
	PasswordHash string
	CreatedAt    time.Time
}
