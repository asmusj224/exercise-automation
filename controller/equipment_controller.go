package controller

import (
	"database/sql"
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type equipmentController struct {
	store services.Store
}

type EquipmentController interface {
	CreateEquipment(ctx *gin.Context)
	GetEquipmentByID(ctx *gin.Context)
	UpdateEquipmentByID(ctx *gin.Context)
	GetEquipments(ctx *gin.Context)
}

func NewEquipmentController(store services.Store) EquipmentController {
	return &equipmentController{
		store: store,
	}
}

type createEquipmentRequest struct {
	Name string `json:"Name" binding:"required"`
}

func (c *equipmentController) CreateEquipment(ctx *gin.Context) {
	var req createEquipmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := sql.NullString{
		String: req.Name,
		Valid:  true,
	}

	equipment, err := c.store.CreateEquipment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, equipment)
}

type getEquipmentRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *equipmentController) GetEquipmentByID(ctx *gin.Context) {
	var req getEquipmentRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equipment, err := c.store.GetEquipmentById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, equipment)
}

type updateEquipmentRequest struct {
	Name string `json:"Name" binding:"required"`
}

func (c *equipmentController) UpdateEquipmentByID(ctx *gin.Context) {
	var req updateEquipmentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateEquipmentByIdParams{
		ID: id,
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
	}

	equipment, err := c.store.UpdateEquipmentById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, equipment)
}

func (c *equipmentController) GetEquipments(ctx *gin.Context) {
	equipment, err := c.store.GetEquipment(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, equipment)
}
