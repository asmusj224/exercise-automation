package controller

import (
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type exerciseEquipmentController struct {
	store services.Store
}

type ExerciseEquipmentController interface {
	CreateExerciseEquipment(ctx *gin.Context)
	GetExerciseEquipmentByID(ctx *gin.Context)
	UpdateExerciseEquipmentByID(ctx *gin.Context)
}

func NewExerciseEquipmentController(store services.Store) ExerciseEquipmentController {
	return &exerciseEquipmentController{
		store: store,
	}
}

type createExerciseEquipmentRequest struct {
	ExerciseID  string `json:"ExerciseID" binding:"required"`
	EquipmentID string `json:"EquipmentID" binding:"required"`
}

func (c *exerciseEquipmentController) CreateExerciseEquipment(ctx *gin.Context) {
	var req createExerciseEquipmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseID, err := uuid.Parse(req.ExerciseID)
	equipmentID, err := uuid.Parse(req.EquipmentID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.CreateExerciseEquipmentParams{
		ExerciseID:  exerciseID,
		EquipmentID: equipmentID,
	}

	exerciseEquipment, err := c.store.CreateExerciseEquipment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseEquipment)
}

type getExerciseEquipmentRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *exerciseEquipmentController) GetExerciseEquipmentByID(ctx *gin.Context) {
	var req getExerciseEquipmentRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseEquipment, err := c.store.GetExerciseEquipmentById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseEquipment)
}

type updateExerciseEquipmentRequest struct {
	ExerciseID  string `json:"ExerciseID" binding:"required"`
	EquipmentID string `json:"EquipmentID" binding:"required"`
}

func (c *exerciseEquipmentController) UpdateExerciseEquipmentByID(ctx *gin.Context) {
	var req updateExerciseEquipmentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	exerciseID, err := uuid.Parse(req.ExerciseID)
	equipmentID, err := uuid.Parse(req.EquipmentID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateExerciseEquipmentByIdParams{
		ID:          id,
		ExerciseID:  exerciseID,
		EquipmentID: equipmentID,
	}

	exerciseEquipment, err := c.store.UpdateExerciseEquipmentById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exerciseEquipment)
}
