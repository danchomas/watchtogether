package transporthttp

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"watchtogether/internal/features/rooms/model"
)

type Rooms interface {
	Create(title, videoURL, name string) (model.Snapshot, string, error)
	Join(code, name string) (model.Snapshot, string, error)
	Get(code, token string) (model.Snapshot, error)
	Subscribe(code, token string) (<-chan model.Event, func(), error)
	UpdatePlayback(code, token string, position float64, playing bool) error
}

type Handler struct{ rooms Rooms }

func New(rooms Rooms) *Handler { return &Handler{rooms: rooms} }

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/rooms", h.create)
	mux.HandleFunc("POST /api/rooms/{code}/join", h.join)
	mux.HandleFunc("GET /api/rooms/{code}", h.get)
	mux.HandleFunc("GET /api/rooms/{code}/events", h.events)
	mux.HandleFunc("POST /api/rooms/{code}/playback", h.playback)
}

func write(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func fail(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, model.ErrInvalid):
		status = http.StatusBadRequest
	case errors.Is(err, model.ErrUnauthorized):
		status = http.StatusForbidden
	case errors.Is(err, model.ErrNotFound):
		status = http.StatusNotFound
	}
	if status == http.StatusInternalServerError {
		write(w, status, map[string]string{"error": "внутренняя ошибка"})
		return
	}
	write(w, status, map[string]string{"error": err.Error()})
}

func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(dst); err != nil {
		fail(w, model.ErrInvalid)
		return false
	}
	var extra any
	if err := d.Decode(&extra); !errors.Is(err, io.EOF) {
		fail(w, model.ErrInvalid)
		return false
	}
	return true
}

func token(r *http.Request) string {
	return strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title    string `json:"title"`
		VideoURL string `json:"videoUrl"`
		Name     string `json:"name"`
	}
	if !decode(w, r, &body) {
		return
	}
	state, key, err := h.rooms.Create(body.Title, body.VideoURL, body.Name)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, map[string]any{"room": state, "token": key})
}

func (h *Handler) join(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if !decode(w, r, &body) {
		return
	}
	state, key, err := h.rooms.Join(r.PathValue("code"), body.Name)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, map[string]any{"room": state, "token": key})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	state, err := h.rooms.Get(r.PathValue("code"), token(r))
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, state)
}
