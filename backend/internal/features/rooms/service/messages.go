package service

import (
	"strings"
	"time"

	"watchtogether/internal/features/rooms/model"
)

func (s *Service) SendMessage(code, token, text string) error {
	text = strings.TrimSpace(text)
	if len([]rune(text)) < 1 || len([]rune(text)) > 500 {
		return model.ErrInvalid
	}
	id, err := randomID(8)
	if err != nil {
		return err
	}
	_, err = s.rooms.Update(normalizeCode(code), func(room *model.Room) error {
		if err := authorize(*room, token); err != nil {
			return err
		}
		room.State.Messages = append(room.State.Messages, model.Message{ID: id, Name: room.Members[token], Text: text, Time: time.Now().UTC()})
		if len(room.State.Messages) > 100 {
			room.State.Messages = room.State.Messages[len(room.State.Messages)-100:]
		}
		return nil
	})
	return err
}
