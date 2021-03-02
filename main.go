package main

import (
	"net/http"

	"github.com/asmusj224/exercise-automation/controller"
	"github.com/asmusj224/exercise-automation/database"
	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	app := gin.Default()

	err := database.Connect()

	if err != nil {
		panic(err)
	}

	store := services.NewStore(database.DB)
	workoutController := controller.NewWorkoutController(store)
	app.POST("/api/v1/workout", workoutController.CreateWorkout)
	app.GET("/api/v1/workout/:id", workoutController.GetWorkoutByID)
	app.GET("/api/v1/workout", workoutController.GetWorkouts)
	app.PUT("/api/v1/workout/:id", workoutController.UpdateWorkoutByID)

	equipmentController := controller.NewEquipmentController(store)
	app.POST("/api/v1/equipment", equipmentController.CreateEquipment)
	app.GET("/api/v1/equipment/:id", equipmentController.GetEquipmentByID)
	app.GET("/api/v1/equipment", equipmentController.GetEquipments)
	app.PUT("/api/v1/equipment/:id", equipmentController.UpdateEquipmentByID)

	exerciseController := controller.NewExerciseController(store)
	app.POST("/api/v1/exercise", exerciseController.CreateExercise)
	app.GET("/api/v1/exercise/:id", exerciseController.GetExerciseByID)
	app.GET("/api/v1/exercise", exerciseController.GetExercises)
	app.PUT("/api/v1/exercise/:id", exerciseController.UpdateExerciseByID)

	muscleController := controller.NewMuscleController(store)
	app.POST("/api/v1/muscle", muscleController.CreateMuscle)
	app.GET("/api/v1/muscle/:id", muscleController.GetMuscleByID)
	app.GET("/api/v1/muscle", muscleController.GetMuscles)
	app.PUT("/api/v1/muscle/:id", muscleController.UpdateMuscleByID)

	exerciseEquipment := controller.NewExerciseEquipmentController(store)
	app.POST("/api/v1/exercise-equipment", exerciseEquipment.CreateExerciseEquipment)
	app.GET("/api/v1/exercise-equipment/:id", exerciseEquipment.GetExerciseEquipmentByID)
	app.PUT("/api/v1/exercise-equipment/:id", exerciseEquipment.UpdateExerciseEquipmentByID)

	exerciseMuscle := controller.NewExerciseMuscleController(store)
	app.POST("/api/v1/exercise-muscle", exerciseMuscle.CreateExerciseMuscle)
	app.GET("/api/v1/exercise-muscle/:id", exerciseMuscle.GetExerciseMuscleByID)
	app.PUT("/api/v1/exercise-muscle/:id", exerciseMuscle.UpdateExerciseMuscleByID)

	exerciseWorkout := controller.NewExerciseWorkoutController(store)
	app.POST("/api/v1/exercise-workout", exerciseWorkout.CreateExerciseWorkout)
	app.GET("/api/v1/exercise-workout/:id", exerciseWorkout.GetExerciseWorkoutByID)
	app.PUT("/api/v1/exercise-workout/:id", exerciseWorkout.UpdateExerciseWorkoutByID)

	app.GET("/test", func(c *gin.Context) {
		id, _ := uuid.Parse("cf89c692-050d-448c-ae57-4ba3bb4ee518")
		found, _ := store.GetExerciseWorkoutByWorkoutId(c, id)
		c.JSON(http.StatusOK, gin.H{
			"result": found,
		})
	})

	defer database.DB.Close()

	app.Run(":3000")
}
