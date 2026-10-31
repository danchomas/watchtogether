package memory_test

import (
	"errors"
	"testing"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
)

func TestRepositoryIsolation(t *testing.T) {
	repo := memory.New()
	room := model.Room{State: model.Snapshot{Code: "ROOM"}, Members: map[string]string{"owner": "Аня"}}
	if err := repo.Insert(room); err != nil {
		t.Fatal(err)
	}
	room.Members["owner"] = "changed"
	stored, err := repo.Get("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Members["owner"] != "Аня" {
		t.Fatal("insert leaked its input")
	}
	stored.Members["owner"] = "changed"
	_, err = repo.Update("ROOM", func(room *model.Room) error { room.Members["owner"] = "changed"; return model.ErrInvalid })
	if !errors.Is(err, model.ErrInvalid) {
		t.Fatalf("unexpected update error: %v", err)
	}
	stored, err = repo.Get("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Members["owner"] != "Аня" {
		t.Fatal("read or failed transaction mutated storage")
	}
	if err := repo.Insert(room); !errors.Is(err, model.ErrConflict) {
		t.Fatalf("expected duplicate room, got %v", err)
	}
}
