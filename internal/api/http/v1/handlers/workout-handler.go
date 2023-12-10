package httpV1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	apiports "github.com/vandaimer/workouts/internal/domain/api-ports"
	"github.com/vandaimer/workouts/internal/domain/services"
)

type WorkoutHandler struct {
	service *services.WorkoutService
}

func NewWorkoutHandler(workoutService *services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		service: workoutService,
	}
}

func (handler *WorkoutHandler) Create(ctx *gin.Context) {
	nweeksString, ok := ctx.GetQuery("nweeks")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query param 'nweeks' is required."})
		return
	}

	nweeks, err := strconv.ParseInt(nweeksString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var workoutRequest apiports.CreateWorkoutRequest
	if err := ctx.ShouldBindJSON(&workoutRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := handler.service.Create(ctx, uint(nweeks), workoutRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}
