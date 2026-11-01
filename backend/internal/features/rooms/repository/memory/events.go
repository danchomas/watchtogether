package memory

import "watchtogether/internal/features/rooms/model"

func (r *Repository) Subscribe(code string) (<-chan model.Event, func(), error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	room, exists := r.rooms[code]
	if !exists {
		return nil, nil, model.ErrNotFound
	}
	ch := make(chan model.Event, 1)
	if r.listeners[code] == nil {
		r.listeners[code] = make(map[chan model.Event]struct{})
	}
	r.listeners[code][ch] = struct{}{}
	ch <- model.Event{Type: "state", Snapshot: room.Clone().State}
	return ch, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		delete(r.listeners[code], ch)
		if len(r.listeners[code]) == 0 {
			delete(r.listeners, code)
		}
	}, nil
}

func (r *Repository) publish(code string, room model.Room) {
	for ch := range r.listeners[code] {
		event := model.Event{Type: "state", Snapshot: room.Clone().State}
		select {
		case ch <- event:
		default:
			// Full snapshots let slow consumers skip stale intermediate states.
			select {
			case <-ch:
			default:
			}
			select {
			case ch <- event:
			default:
			}
		}
	}
}
