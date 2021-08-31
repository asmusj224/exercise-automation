package main

import (
	"log"

	"github.com/asmusj224/exercise-automation/config"
	"github.com/asmusj224/exercise-automation/controller"
	"github.com/asmusj224/exercise-automation/database"
	"github.com/asmusj224/exercise-automation/services"
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

	defer database.DB.Close()

	app.Run(":" + port)
}
