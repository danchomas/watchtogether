package model

import "errors"

var (
	ErrNotFound     = errors.New("комната не найдена")
	ErrConflict     = errors.New("код комнаты уже используется")
	ErrUnauthorized = errors.New("сначала войдите в комнату")
	ErrInvalid      = errors.New("проверьте введённые данные")
)

type Snapshot struct {
	Code         string   `json:"code"`
	Title        string   `json:"title"`
	VideoURL     string   `json:"videoUrl"`
	Participants int      `json:"participants"`
	Playback     Playback `json:"playback"`
}

type Room struct {
	State   Snapshot
	Members map[string]string
}

func (r Room) Clone() Room {
	copy := r
	copy.Members = make(map[string]string, len(r.Members))
	for token, name := range r.Members {
		copy.Members[token] = name
	}
	return copy
}
