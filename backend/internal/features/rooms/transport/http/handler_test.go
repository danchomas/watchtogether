package transporthttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"watchtogether/internal/features/rooms/model"
	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
)

func TestHTTPRoomFlow(t *testing.T) {
	mux := http.NewServeMux()
	New(service.New(memory.New())).Register(mux)
	request := func(method, path, body, token string) *httptest.ResponseRecorder {
		t.Helper()
		r := httptest.NewRequest(method, path, bytes.NewBufferString(body))
		if token != "" {
			r.Header.Set("Authorization", "Bearer "+token)
		}
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)
		return w
	}
	created := request("POST", "/api/rooms", `{"title":"Вечер","videoUrl":"https://example.org/video.mp4","name":"Аня"}`, "")
	if created.Code != http.StatusCreated {
		t.Fatal(created.Body.String())
	}
	var result struct {
		Room  model.Snapshot `json:"room"`
		Token string         `json:"token"`
	}
	if err := json.Unmarshal(created.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	path := "/api/rooms/" + result.Room.Code
	if got := request("GET", path, "", "").Code; got != http.StatusForbidden {
		t.Fatalf("unauthenticated GET: %d", got)
	}
	joined := request("POST", path+"/join", `{"name":"Борис"}`, "")
	if joined.Code != http.StatusOK {
		t.Fatal(joined.Body.String())
	}
	if got := request("POST", path+"/playback", `{"position":10,"playing":true}`, result.Token).Code; got != http.StatusNoContent {
		t.Fatalf("playback: %d", got)
	}
	if got := request("POST", path+"/messages", `{"text":"Привет"}`, result.Token).Code; got != http.StatusNoContent {
		t.Fatalf("message: %d", got)
	}
	state := request("GET", path, "", result.Token)
	if state.Code != http.StatusOK {
		t.Fatal(state.Body.String())
	}
	var room model.Snapshot
	if err := json.Unmarshal(state.Body.Bytes(), &room); err != nil {
		t.Fatal(err)
	}
	if room.Participants != 2 || !room.Playback.Playing || room.Playback.Position != 10 || len(room.Messages) != 1 || room.Messages[0].Name != "Аня" {
		t.Fatalf("unexpected state: %+v", room)
	}
}
