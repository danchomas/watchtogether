package service

import (
	"math"
	"time"

	"watchtogether/internal/features/rooms/model"
)

func (s *Service) Subscribe(code, token string) (<-chan model.Event, func(), error) {
	if _, err := s.Get(code, token); err != nil {
		return nil, nil, err
	}
	return s.rooms.Subscribe(normalizeCode(code))
}

func (s *Service) UpdatePlayback(code, token string, position float64, playing bool) error {
	if math.IsNaN(position) || math.IsInf(position, 0) || position < 0 || position > 86400 {
		return model.ErrInvalid
	}
	_, err := s.rooms.Update(normalizeCode(code), func(room *model.Room) error {
		if err := authorize(*room, token); err != nil {
			return err
		}
		room.State.Playback = model.Playback{Position: position, Playing: playing, UpdatedAt: time.Now().UTC()}
		return nil
	})
	return err
}
