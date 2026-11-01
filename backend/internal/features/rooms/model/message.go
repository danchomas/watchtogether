package model

import "time"

type Message struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}
