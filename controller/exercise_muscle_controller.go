package controller

import (
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type exerciseMuscleController struct {
	store services.Store
}

type ExerciseMuscleController interface {
	CreateExerciseMuscle(ctx *gin.Context)
	GetExerciseMuscleByID(ctx *gin.Context)
	UpdateExerciseMuscleByID(ctx *gin.Context)
}

func NewExerciseMuscleController(store services.Store) ExerciseMuscleController {
	return &exerciseMuscleController{
		store: store,
	}
}

type createExerciseMuscleRequest struct {
	ExerciseID string `json:"ExerciseID" binding:"required"`
	MuscleID   string `json:"MuscleID" binding:"required"`
}

func (c *exerciseMuscleController) CreateExerciseMuscle(ctx *gin.Context) {
	var req createExerciseMuscleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseID, err := uuid.Parse(req.ExerciseID)
	muscleID, err := uuid.Parse(req.MuscleID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.CreateExerciseMuscleParams{
		ExerciseID: exerciseID,
		MuscleID:   muscleID,
	}

	exerciseMuscle, err := c.store.CreateExerciseMuscle(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseMuscle)
}

type getExerciseMuscleRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *exerciseMuscleController) GetExerciseMuscleByID(ctx *gin.Context) {
	var req getExerciseMuscleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseMuscle, err := c.store.GetExerciseMuscleById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseMuscle)
}

type updateExerciseMuscleRequest struct {
	ExerciseID string `json:"ExerciseID" binding:"required"`
	MuscleID   string `json:"MuscleID" binding:"required"`
}

func (c *exerciseMuscleController) UpdateExerciseMuscleByID(ctx *gin.Context) {
	var req updateExerciseMuscleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	exerciseID, err := uuid.Parse(req.ExerciseID)
	muscleID, err := uuid.Parse(req.MuscleID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateExerciseMuscleByIdParams{
		ID:         id,
		ExerciseID: exerciseID,
		MuscleID:   muscleID,
	}

	exerciseMuscle, err := c.store.UpdateExerciseMuscleById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseMuscle)
}
