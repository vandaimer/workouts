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

	_, err := service.repository.Create(ctx, spiWorkouts)
	if err != nil {
		return nil, err
	}

	response := model.WorkoutResponse{
		MediumDistance:       8000,
		MediumTime:           3000,
		MaxDistance:          22000,
		MaxTime:              7000,
		MediumWeeklyDistance: 30000,
		MediumWeeklyTime:     9000,
		MaxWeeklyDistance:    30000,
		MaxWeeklyTime:        9000,
	}

	return &response, nil
}
