package model

import "time"

// Unit of work assigned to a user
type Task struct {
	ID             string
	UserID         string
	Title          string
	Description    string
	Priority       float64
	WorkedHours    float64
	AvailableHours float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DoneAt         *time.Time // nil if not compleated
}
