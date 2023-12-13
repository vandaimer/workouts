package spiports

import (
	"context"
	"time"

	"github.com/vandaimer/workouts/internal/domain/model"
)

type CreateWorkoutRequest struct {
	WeekNumber uint
	Distance   int
	Time       int
	Timestamp  time.Time
}

type WorkoutRepository interface {
	Create(ctx context.Context, request []CreateWorkoutRequest) ([]model.Workout, error)
}
