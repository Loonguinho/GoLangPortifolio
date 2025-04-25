package models

import "time"

type Task struct {
	ID          string
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}
