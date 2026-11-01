package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/url"
	"strings"
	"time"

	"watchtogether/internal/features/rooms/model"
)

type Repository interface {
	Insert(room model.Room) error
	Get(code string) (model.Room, error)
	Update(code string, change func(*model.Room) error) (model.Room, error)
	Subscribe(code string) (<-chan model.Event, func(), error)
}

type Service struct{ rooms Repository }

func New(rooms Repository) *Service { return &Service{rooms: rooms} }

func randomID(size int) (string, error) {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func validName(name string) bool { return len([]rune(name)) >= 1 && len([]rune(name)) <= 40 }

func normalizeCode(code string) string { return strings.ToUpper(strings.TrimSpace(code)) }

func authorize(room model.Room, token string) error {
	if _, ok := room.Members[token]; !ok {
		return model.ErrUnauthorized
	}
	return nil
}

func (s *Service) Create(title, videoURL, name string) (model.Snapshot, string, error) {
	title, name, videoURL = strings.TrimSpace(title), strings.TrimSpace(name), strings.TrimSpace(videoURL)
	u, err := url.Parse(videoURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.User != nil || len(videoURL) > 2048 || len([]rune(title)) < 1 || len([]rune(title)) > 100 || !validName(name) {
		return model.Snapshot{}, "", model.ErrInvalid
	}
	token, err := randomID(24)
	if err != nil {
		return model.Snapshot{}, "", err
	}
	for {
		code, err := randomID(4)
		if err != nil {
			return model.Snapshot{}, "", err
		}
		room := model.Room{State: model.Snapshot{Code: normalizeCode(code), Title: title, VideoURL: videoURL, Participants: 1}, Members: map[string]string{token: name}}
		room.State.Playback = model.Playback{UpdatedAt: time.Now().UTC()}
		room.State.Messages = []model.Message{}
		err = s.rooms.Insert(room)
		if errors.Is(err, model.ErrConflict) {
			continue
		}
		if err != nil {
			return model.Snapshot{}, "", err
		}
		return room.State, token, nil
	}
}

func (s *Service) Join(code, name string) (model.Snapshot, string, error) {
	name = strings.TrimSpace(name)
	if !validName(name) {
		return model.Snapshot{}, "", model.ErrInvalid
	}
	token, err := randomID(24)
	if err != nil {
		return model.Snapshot{}, "", err
	}
	room, err := s.rooms.Update(normalizeCode(code), func(room *model.Room) error {
		room.Members[token] = name
		room.State.Participants = len(room.Members)
		return nil
	})
	if err != nil {
		return model.Snapshot{}, "", err
	}
	return room.State, token, nil
}

func (s *Service) Get(code, token string) (model.Snapshot, error) {
	room, err := s.rooms.Get(normalizeCode(code))
	if err != nil {
		return model.Snapshot{}, err
	}
	if err := authorize(room, token); err != nil {
		return model.Snapshot{}, err
	}
	return room.State, nil
}
