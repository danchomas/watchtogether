package service_test

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
)

func TestMessages(t *testing.T) {
	s := service.New(memory.New())
	room, owner, err := s.Create("Кино", "https://example.org/movie.mp4", "Аня")
	if err != nil {
		t.Fatal(err)
	}
	_, guest, err := s.Join(room.Code, "Борис")
	if err != nil {
		t.Fatal(err)
	}
	if err := s.SendMessage(room.Code, "wrong", "Привет"); !errors.Is(err, model.ErrUnauthorized) {
		t.Fatalf("expected unauthorized, got %v", err)
	}
	for _, text := range []string{" ", strings.Repeat("я", 501)} {
		if err := s.SendMessage(room.Code, guest, text); !errors.Is(err, model.ErrInvalid) {
			t.Fatalf("expected invalid text, got %v", err)
		}
	}
	for i := 0; i < 105; i++ {
		if err := s.SendMessage(room.Code, guest, fmt.Sprintf("Сообщение %d", i)); err != nil {
			t.Fatal(err)
		}
	}
	state, err := s.Get(room.Code, owner)
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Messages) != 100 || state.Messages[0].Text != "Сообщение 5" || state.Messages[0].Name != "Борис" {
		t.Fatalf("unexpected history: %+v", state.Messages)
	}
	state.Messages[0].Text = "changed"
	stored, err := s.Get(room.Code, owner)
	if err != nil {
		t.Fatal(err)
	}
	if stored.Messages[0].Text != "Сообщение 5" {
		t.Fatal("snapshot mutated repository")
	}
}

func TestConcurrentMessages(t *testing.T) {
	s := service.New(memory.New())
	room, token, err := s.Create("Кино", "https://example.org/movie.mp4", "Аня")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.SendMessage(room.Code, token, "Привет"); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	state, err := s.Get(room.Code, token)
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Messages) != 50 {
		t.Fatalf("lost concurrent messages: %d", len(state.Messages))
	}
}
