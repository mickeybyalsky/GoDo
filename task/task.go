package task

import (
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Created   time.Time
	Completed bool
	CompletedAt *time.Time
}
