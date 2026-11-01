package memory_test

import (
	"testing"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
)

func TestMessageStorageAndEventIsolation(t *testing.T) {
	repo := memory.New()
	input := model.Room{State: model.Snapshot{Code: "ROOM", Messages: []model.Message{{Text: "Привет"}}}}
	if err := repo.Insert(input); err != nil {
		t.Fatal(err)
	}
	input.State.Messages[0].Text = "changed"
	stream, cancel, err := repo.Subscribe("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()
	event := <-stream
	if event.Snapshot.Messages[0].Text != "Привет" {
		t.Fatal("insert did not isolate history")
	}
	event.Snapshot.Messages[0].Text = "changed"
	stored, err := repo.Get("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	if stored.State.Messages[0].Text != "Привет" {
		t.Fatal("event mutated stored history")
	}
	if _, err := repo.Update("ROOM", func(room *model.Room) error {
		room.State.Messages = append(room.State.Messages, model.Message{Text: "Пока"})
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	event = <-stream
	if len(event.Snapshot.Messages) != 2 {
		t.Fatal("missing message update")
	}
	event.Snapshot.Messages[1].Text = "changed"
	stored, err = repo.Get("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	if stored.State.Messages[1].Text != "Пока" {
		t.Fatal("published state leaked storage")
	}
}
