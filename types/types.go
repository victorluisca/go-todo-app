package types

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
}
