package model

import "time"

type Playback struct {
	Position  float64   `json:"position"`
	Playing   bool      `json:"playing"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Event struct {
	Type     string   `json:"type"`
	Snapshot Snapshot `json:"snapshot"`
}
