package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type exerciseWorkoutController struct {
	store services.Store
}

type ExerciseWorkoutController interface {
	CreateExerciseWorkout(ctx *gin.Context)
	GetExerciseWorkoutByID(ctx *gin.Context)
	UpdateExerciseWorkoutByID(ctx *gin.Context)
	EmailRandomExerciseWorkout(ctx *gin.Context)
}

func NewExerciseWorkoutController(store services.Store) ExerciseWorkoutController {
	return &exerciseWorkoutController{
		store: store,
	}
}

type createExerciseWorkoutRequest struct {
	ExerciseID string `json:"ExerciseID" binding:"required"`
	WorkoutID  string `json:"WorkoutID" binding:"required"`
}

func (c *exerciseWorkoutController) CreateExerciseWorkout(ctx *gin.Context) {
	var req createExerciseWorkoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseID, err := uuid.Parse(req.ExerciseID)
	WorkoutID, err := uuid.Parse(req.WorkoutID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.CreateExerciseWorkoutParams{
		ExerciseID: exerciseID,
		WorkoutID:  WorkoutID,
	}

	exerciseWorkout, err := c.store.CreateExerciseWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseWorkout)
}

type getExerciseWorkoutRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *exerciseWorkoutController) GetExerciseWorkoutByID(ctx *gin.Context) {
	var req getExerciseWorkoutRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseWorkout, err := c.store.GetExerciseWorkoutById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseWorkout)
}

type updateExerciseWorkoutRequest struct {
	ExerciseID string `json:"ExerciseID" binding:"required"`
	WorkoutID  string `json:"WorkoutID" binding:"required"`
}

func (c *exerciseWorkoutController) UpdateExerciseWorkoutByID(ctx *gin.Context) {
	var req updateExerciseWorkoutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	exerciseID, err := uuid.Parse(req.ExerciseID)
	WorkoutID, err := uuid.Parse(req.WorkoutID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateExerciseWorkoutByIdParams{
		ID:         id,
		ExerciseID: exerciseID,
		WorkoutID:  WorkoutID,
	}

	exerciseWorkout, err := c.store.UpdateExerciseWorkoutById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseWorkout)
}

func (c *exerciseWorkoutController) EmailRandomExerciseWorkout(ctx *gin.Context) {

	emailService := services.NewEmailService()
	workout, _ := c.store.GetRandomExerciseWorkout(ctx)
	var exercises []services.Exercise
	json.Unmarshal(workout.Exercises, &exercises)
	subject := "Exercise Plan"
	email_body := ""
	for _, exercise := range exercises {
		email_body += exercise.Name + " number of reps: " + strconv.Itoa(int(exercise.NumberOfReps)) + " number of sets: " + strconv.Itoa(int(exercise.NumberOfSets)) + "\n"
	}
	_, err := emailService.SendEmail("jeffrey.asmus88@gmail.com", subject, email_body)
	if err != nil {
		log.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, workout)
}
