package memory

import (
	"sync"

	"watchtogether/internal/features/rooms/model"
)

type Repository struct {
	mu    sync.Mutex
	rooms map[string]model.Room
}

func New() *Repository { return &Repository{rooms: make(map[string]model.Room)} }

func (r *Repository) Insert(room model.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.rooms[room.State.Code]; exists {
		return model.ErrConflict
	}
	r.rooms[room.State.Code] = room.Clone()
	return nil
}

func (r *Repository) Get(code string) (model.Room, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	room, exists := r.rooms[code]
	if !exists {
		return model.Room{}, model.ErrNotFound
	}
	return room.Clone(), nil
}

// Update applies a transaction to an isolated copy; errors leave the stored room unchanged.
func (r *Repository) Update(code string, change func(*model.Room) error) (model.Room, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	room, exists := r.rooms[code]
	if !exists {
		return model.Room{}, model.ErrNotFound
	}
	room = room.Clone()
	if err := change(&room); err != nil {
		return model.Room{}, err
	}
	r.rooms[code] = room.Clone()
	return room, nil
}
