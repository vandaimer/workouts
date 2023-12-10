package apiports

import "time"

type CreateWorkoutRequest []struct {
	WeekNumber int       `json:"weekNumber"`
	Distance   int       `json:"distance"`
	Time       int       `json:"time"`
	Timestamp  time.Time `json:"timestamp"`
}
