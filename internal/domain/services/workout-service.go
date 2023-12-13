package services

import (
	"context"

	apiports "github.com/vandaimer/workouts/internal/domain/api-ports"
	"github.com/vandaimer/workouts/internal/domain/model"
	spiports "github.com/vandaimer/workouts/internal/domain/spi-ports"
)

type WorkoutService struct {
	repository spiports.WorkoutRepository
}

func NewWorkoutService(repository spiports.WorkoutRepository) *WorkoutService {
	return &WorkoutService{
		repository: repository,
	}
}

func (service *WorkoutService) Create(ctx context.Context, nweek uint, request apiports.CreateWorkoutRequest) (*model.WorkoutResponse, error) {
	spiWorkouts := make([]spiports.CreateWorkoutRequest, 0)

	for _, v := range request {
		spiWorkouts = append(spiWorkouts, spiports.CreateWorkoutRequest{
			WeekNumber: nweek,
			Distance:   v.Distance,
			Time:       v.Time,
			Timestamp:  v.Timestamp,
		})
	}

	workoutsCreated, err := service.repository.Create(ctx, spiWorkouts)
	if err != nil {
		return nil, err
	}

	return workoutsCreated, nil
}
