package model

import "time"

type Workout struct {
	Id         string    `json:"id"`
	WeekNumber uint      `json:"week_name"`
	Distance   int       `json:"distance"`
	Time       int       `json:"time"`
	Timestamp  time.Time `json:"timestamp"`
}

type WorkoutResponse struct {
	MediumDistance       int `json:"medium_distance"`
	MediumTime           int `json:"medium_time"`
	MaxDistance          int `json:"max_distance"`
	MaxTime              int `json:"max_time"`
	MediumWeeklyDistance int `json:"medium_weekly_distance"`
	MediumWeeklyTime     int `json:"medium_weekly_time"`
	MaxWeeklyDistance    int `json:"max_weekly_distance"`
	MaxWeeklyTime        int `json:"max_weekly_time"`
}
