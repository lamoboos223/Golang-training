package domain

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
