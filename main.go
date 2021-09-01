package main

import (
	"log"
	"net/http"

	"github.com/asmusj224/exercise-automation/config"
	"github.com/asmusj224/exercise-automation/controller"
	"github.com/asmusj224/exercise-automation/database"
	"github.com/asmusj224/exercise-automation/services"
	"github.com/asmusj224/exercise-automation/workers"
	"github.com/gin-gonic/gin"
)

func main() {

	port := config.Config("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := gin.Default()

	err := database.Connect()

	if err != nil {
		panic(err)
	}

	app.GET("/api/v1/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "good"})
	})

	store := services.NewStore(database.DB)
	workoutController := controller.NewWorkoutController(store)
	app.POST("/api/v1/workout", workoutController.CreateWorkout)
	app.GET("/api/v1/workout/:id", workoutController.GetWorkoutByID)
	app.GET("/api/v1/workout", workoutController.GetWorkouts)
	app.PUT("/api/v1/workout/:id", workoutController.UpdateWorkoutByID)

	exerciseController := controller.NewExerciseController(store)
	app.POST("/api/v1/exercise", exerciseController.CreateExercise)
	app.GET("/api/v1/exercise/:id", exerciseController.GetExerciseByID)
	app.GET("/api/v1/exercise", exerciseController.GetExercises)
	app.PUT("/api/v1/exercise/:id", exerciseController.UpdateExerciseByID)

	exerciseWorkout := controller.NewExerciseWorkoutController(store)
	app.POST("/api/v1/exercise-workout", exerciseWorkout.CreateExerciseWorkout)
	app.GET("/api/v1/exercise-workout/:id", exerciseWorkout.GetExerciseWorkoutByID)
	app.PUT("/api/v1/exercise-workout/:id", exerciseWorkout.UpdateExerciseWorkoutByID)

	workers.NewWorkers().Start()
	defer database.DB.Close()

	app.Run(":" + port)
}
