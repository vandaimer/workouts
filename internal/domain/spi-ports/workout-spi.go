package spiports

import (
	"context"
	"time"

	"github.com/vandaimer/workouts/internal/domain/model"
)

type CreateWorkoutRequest struct {
	WeekNumber uint      `json:"week_name"`
	Distance   int       `json:"distance"`
	Time       int       `json:"time"`
	Timestamp  time.Time `json:"timestamp"`
}

type WorkoutRepository interface {
	Create(ctx context.Context, request []CreateWorkoutRequest) ([]model.Workout, error)
}
