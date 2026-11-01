package transporthttp

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"net/http/httptest"
	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
)

func TestEventStream(t *testing.T) {
	s := service.New(memory.New())
	room, token, err := s.Create("Кино", "https://example.org/movie.mp4", "Аня")
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	New(s).Register(mux)
	server := httptest.NewServer(mux)
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", server.URL+"/api/rooms/"+room.Code+"/events?token="+token, nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 || response.Header.Get("Content-Type") != "text/event-stream" {
		t.Fatalf("unexpected stream response: %+v", response)
	}
	scanner := bufio.NewScanner(response.Body)
	readState := func() model.Event {
		t.Helper()
		for scanner.Scan() {
			if data, ok := strings.CutPrefix(scanner.Text(), "data: "); ok {
				var event model.Event
				if err := json.Unmarshal([]byte(data), &event); err != nil {
					t.Fatal(err)
				}
				return event
			}
		}
		t.Fatalf("stream closed: %v", scanner.Err())
		return model.Event{}
	}
	if initial := readState(); initial.Snapshot.Code != room.Code {
		t.Fatal("wrong initial state")
	}
	if err := s.UpdatePlayback(room.Code, token, 20, true); err != nil {
		t.Fatal(err)
	}
	if next := readState(); next.Snapshot.Playback.Position != 20 || !next.Snapshot.Playback.Playing {
		t.Fatalf("unexpected playback: %+v", next)
	}
}
