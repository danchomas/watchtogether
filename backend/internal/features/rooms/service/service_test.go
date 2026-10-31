package service_test

import (
	"errors"
	"strings"
	"testing"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
)

func TestRoomFlow(t *testing.T) {
	s := service.New(memory.New())
	room, owner, err := s.Create(" Кино ", "https://example.org/movie.mp4", " Аня ")
	if err != nil {
		t.Fatal(err)
	}
	if room.Title != "Кино" || room.Participants != 1 || len(room.Code) != 8 {
		t.Fatalf("unexpected room: %+v", room)
	}
	_, guest, err := s.Join(strings.ToLower(room.Code), "Борис")
	if err != nil {
		t.Fatal(err)
	}
	if guest == owner {
		t.Fatal("tokens must be unique")
	}
	state, err := s.Get(room.Code, guest)
	if err != nil {
		t.Fatal(err)
	}
	if state.Participants != 2 {
		t.Fatalf("unexpected participant count: %d", state.Participants)
	}
	if _, err := s.Get(room.Code, "wrong"); !errors.Is(err, model.ErrUnauthorized) {
		t.Fatalf("expected unauthorized, got %v", err)
	}
}

func TestInvalidInput(t *testing.T) {
	s := service.New(memory.New())
	if _, _, err := s.Create("Кино", "javascript:alert(1)", "Аня"); !errors.Is(err, model.ErrInvalid) {
		t.Fatalf("expected invalid URL, got %v", err)
	}
	if _, _, err := s.Join("missing", " "); !errors.Is(err, model.ErrInvalid) {
		t.Fatalf("expected invalid name, got %v", err)
	}
	if _, _, err := s.Join("missing", "Аня"); !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("expected missing room, got %v", err)
	}
}
