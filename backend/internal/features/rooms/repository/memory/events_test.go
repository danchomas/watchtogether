package memory_test

import (
	"testing"
	"time"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
)

func TestSubscriptionKeepsLatestSnapshot(t *testing.T) {
	repo := memory.New()
	if err := repo.Insert(model.Room{State: model.Snapshot{Code: "ROOM"}}); err != nil {
		t.Fatal(err)
	}
	stream, cancel, err := repo.Subscribe("ROOM")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()
	for i := 1; i <= 10; i++ {
		_, err := repo.Update("ROOM", func(room *model.Room) error { room.State.Playback.Position = float64(i); return nil })
		if err != nil {
			t.Fatal(err)
		}
	}
	select {
	case event := <-stream:
		if event.Snapshot.Playback.Position != 10 {
			t.Fatalf("stale snapshot: %+v", event)
		}
	case <-time.After(time.Second):
		t.Fatal("missing state")
	}
	cancel()
	cancel() // cleanup must be idempotent
}
