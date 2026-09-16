package room

import (
	"time"
	"watchroom/internal/domain/models/video"
)

type Room struct {
	Id           string
	Title        string
	Code         string
	Visibility   Visibility
	PasswordHash string
	CreatedAt    time.Time
	Video        video.Video
}
