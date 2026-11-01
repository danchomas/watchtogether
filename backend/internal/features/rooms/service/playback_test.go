package service_test

import (
	"errors"
	"math"
	"testing"
	"time"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
)

func TestPlayback(t *testing.T) {
	s := service.New(memory.New())
	room, token, err := s.Create("Кино", "https://example.org/movie.mp4", "Аня")
	if err != nil {
		t.Fatal(err)
	}
	stream, cancel, err := s.Subscribe(room.Code, token)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()
	<-stream
	if err := s.UpdatePlayback(room.Code, token, 12.5, true); err != nil {
		t.Fatal(err)
	}
	select {
	case event := <-stream:
		if !event.Snapshot.Playback.Playing || event.Snapshot.Playback.Position != 12.5 {
			t.Fatalf("unexpected event: %+v", event)
		}
	case <-time.After(time.Second):
		t.Fatal("missing playback event")
	}
	if err := s.UpdatePlayback(room.Code, "wrong", 0, false); !errors.Is(err, model.ErrUnauthorized) {
		t.Fatalf("expected unauthorized, got %v", err)
	}
	for _, position := range []float64{-1, 86401, math.NaN(), math.Inf(1)} {
		if err := s.UpdatePlayback(room.Code, token, position, false); !errors.Is(err, model.ErrInvalid) {
			t.Fatalf("expected invalid position, got %v", err)
		}
	}
	if _, _, err := s.Subscribe(room.Code, "wrong"); !errors.Is(err, model.ErrUnauthorized) {
		t.Fatalf("expected unauthorized subscription, got %v", err)
	}
}
