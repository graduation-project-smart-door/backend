package domain

import "time"

type Event struct {
	ID        int       `json:"id"`
	EventType string    `json:"event_type"`
	Direction string    `json:"direction"`
	UserID    int       `json:"user_id"`
	EventTime time.Time `json:"event_time"`
}
