package transporthttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) playback(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Position float64 `json:"position"`
		Playing  bool    `json:"playing"`
	}
	if !decode(w, r, &body) {
		return
	}
	if err := h.rooms.UpdatePlayback(r.PathValue("code"), token(r), body.Position, body.Playing); err != nil {
		fail(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) events(w http.ResponseWriter, r *http.Request) {
	// EventSource cannot attach an Authorization header; this token is room-scoped.
	ch, unsubscribe, err := h.rooms.Subscribe(r.PathValue("code"), r.URL.Query().Get("token"))
	if err != nil {
		fail(w, err)
		return
	}
	defer unsubscribe()
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unavailable", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no")
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case event := <-ch:
			data, _ := json.Marshal(event)
			if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
				return
			}
			flusher.Flush()
		case <-ticker.C:
			if _, err := fmt.Fprint(w, ": keepalive\n\n"); err != nil {
				return
			}
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
