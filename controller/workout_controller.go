package controller

import (
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type workoutController struct {
	store services.Store
}

type WorkoutController interface {
	CreateWorkout(ctx *gin.Context)
	GetWorkoutByID(ctx *gin.Context)
	UpdateWorkoutByID(ctx *gin.Context)
	GetWorkouts(ctx *gin.Context)
}

func NewWorkoutController(store services.Store) WorkoutController {
	return &workoutController{
		store: store,
	}
}

type createWorkoutRequest struct {
	Name  string `json:"Name" binding:"required"`
	Split string `json:"Split" binding:"required"`
}

func (c *workoutController) CreateWorkout(ctx *gin.Context) {
	var req createWorkoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	arg := services.CreateWorkoutParams{
		Name:  req.Name,
		Split: services.WorkoutSplit(req.Split),
	}
	workout, err := c.store.CreateWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workout)
}

type getWorkoutRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *workoutController) GetWorkoutByID(ctx *gin.Context) {
	var req getWorkoutRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workout, err := c.store.GetWorkoutById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workout)
}

type updateWorkoutRequest struct {
	Name  string `json:"Name" binding:"required"`
	Split string `json:"Split" binding:"required"`
}

func (c *workoutController) UpdateWorkoutByID(ctx *gin.Context) {
	var req updateWorkoutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateWorkoutByIdParams{
		ID:    id,
		Name:  req.Name,
		Split: services.WorkoutSplit(req.Split),
	}

	workout, err := c.store.UpdateWorkoutById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workout)
}

func (c *workoutController) GetWorkouts(ctx *gin.Context) {
	workouts, err := c.store.GetWorkout(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workouts)
}
