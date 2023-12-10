package postgres

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/vandaimer/workouts/internal/domain/model"
	spiports "github.com/vandaimer/workouts/internal/domain/spi-ports"
)

const (
	tableName = "workout"
)

type PostgresWorkoutRepository struct {
	client *BunPostgresDatabaseClient
}

func NewWorkoutRepository(dbClient *BunPostgresDatabaseClient) *PostgresWorkoutRepository {
	log.Info().Msg("Workout Repository Initialized.")
	return &PostgresWorkoutRepository{
		client: dbClient,
	}
}

func (repo *PostgresWorkoutRepository) Create(ctx context.Context, createWorkoutRequest []spiports.CreateWorkoutRequest) (*model.Workout, error) {
	log.Info().Msg("Workout Repository Creating...")
	workout := &model.Workout{}

	_, err := repo.client.DB.NewInsert().Model(&createWorkoutRequest).ModelTableExpr(tableName).Returning("*").Exec(ctx, workout)
	if err != nil {
		return nil, err
	}

	log.Info().Msg(fmt.Sprintf("Workout Repository Created '%s'", workout.Id))

	return workout, nil
}
