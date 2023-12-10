package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/vandaimer/workouts/config"
	"github.com/vandaimer/workouts/internal/api"
	v1Handlers "github.com/vandaimer/workouts/internal/api/http/v1/handlers"
	v1Routes "github.com/vandaimer/workouts/internal/api/http/v1/routes"
	"github.com/vandaimer/workouts/internal/domain/services"
	"github.com/vandaimer/workouts/internal/spi/repositories/postgres"
)

func main() {
	_, cancel := context.WithCancel(context.Background())

	config := config.MustLoad()

	log.Info().Msg("Configuration loaded.")

	dbClient := postgres.NewBunPostgresDatabaseClient(&config.Db)

	workoutRepository := postgres.NewWorkoutRepository(dbClient)

	router := gin.New()
	router.Use(gin.Recovery())

	workoutService := services.NewWorkoutService(workoutRepository)

	workoutHandler := v1Handlers.NewWorkoutHandler(workoutService)

	v1Routes.AttachV1WorkoutRoutes(router, workoutHandler)

	server := api.NewHTTPServer(&config.App, router)

	go server.Run()

	log.Info().Msg("Application successfully initialized.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
	cancel()
	os.Exit(0)
}
