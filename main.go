package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asmusj224/exercise-automation/controller"
	"github.com/asmusj224/exercise-automation/database"
	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	err := database.Connect()

	if err != nil {
		panic(err)
	}

	store := services.NewStore(database.DB)
	messageService := services.NewMessageService(&http.Client{})
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
		found, _ := store.GetRandomExerciseWorkout(c)
		var m []services.Exercise
		json.Unmarshal(found.Exercises, &m)
		str := "______________" + string(found.Name) + "______________" + "\n"
		sid, err := messageService.SendSMS("+12532034540", str)
		for _, v := range m {
			message := v.Name + " " + string(v.NumberOfReps.Int32) + " " + string(v.NumberOfReps.Int32) + "\n"
			messageService.SendSMS("+12532034540", message)
		}
		log.Println(str)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": found,
			"sid":    sid,
		})
	})

	defer database.DB.Close()

	app.Run(":3000")
}
