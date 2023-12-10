package model

import "time"

type Workout struct {
	Id         string    `json:"id"`
	WeekNumber uint      `json:"week_name"`
	Distance   int       `json:"distance"`
	Time       int       `json:"time"`
	Timestamp  time.Time `json:"timestamp"`
}
