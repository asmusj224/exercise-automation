package controller

import (
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type muscleController struct {
	store services.Store
}

type MuscleController interface {
	CreateMuscle(ctx *gin.Context)
	GetMuscleByID(ctx *gin.Context)
	UpdateMuscleByID(ctx *gin.Context)
	GetMuscles(ctx *gin.Context)
}

func NewMuscleController(store services.Store) MuscleController {
	return &muscleController{
		store: store,
	}
}

type createMuscleRequest struct {
	Name string `json:"Name" binding:"required"`
}

func (c *muscleController) CreateMuscle(ctx *gin.Context) {
	var req createMuscleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	muscle, err := c.store.CreateMuscle(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, muscle)
}

type getMuscleRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *muscleController) GetMuscleByID(ctx *gin.Context) {
	var req getMuscleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	muscle, err := c.store.GetMuscleById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, muscle)
}

type updateMuscleRequest struct {
	Name string `json:"Name" binding:"required"`
}

func (c *muscleController) UpdateMuscleByID(ctx *gin.Context) {
	var req updateMuscleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateMuscleByIdParams{
		ID:   id,
		Name: req.Name,
	}

	muscle, err := c.store.UpdateMuscleById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, muscle)
}

func (c *muscleController) GetMuscles(ctx *gin.Context) {
	muscles, err := c.store.GetMuscles(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, muscles)
}
