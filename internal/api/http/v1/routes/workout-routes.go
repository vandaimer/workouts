package httpV1

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/vandaimer/workouts/internal/api/http/v1/handlers"
)

func AttachV1WorkoutRoutes(router *gin.Engine, workoutHandler *handlers.WorkoutHandler) *gin.RouterGroup {
	v1 := router.Group("/api/v1/")
	v1.POST("/analyse", workoutHandler.Create)
	return v1
}
