package controller

import (
	"database/sql"
	"net/http"

	"github.com/asmusj224/exercise-automation/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type exerciseController struct {
	store services.Store
}

type ExerciseController interface {
	CreateExercise(ctx *gin.Context)
	GetExerciseByID(ctx *gin.Context)
	UpdateExerciseByID(ctx *gin.Context)
	GetExercises(ctx *gin.Context)
}

func NewExerciseController(store services.Store) ExerciseController {
	return &exerciseController{
		store: store,
	}
}

type createExerciseRequest struct {
	Name           string `json:"Name" binding:"required"`
	Category       string `json:"category" binding:"required"`
	NumberOfReps   int32  `json:"NumberOfReps"`
	NumberOfSets   int32  `json:"NumberOfSets"`
	RepetitionUnit string `json:"RepetitionUnit"`
	PhotoURL       string `json:"PhotoURL"`
	VideoURL       string `json:"VideoURL"`
}

func (c *exerciseController) CreateExercise(ctx *gin.Context) {
	var req createExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.CreateExerciseParams{
		Name:     req.Name,
		Category: services.ExerciseCategory(req.Category),
		NumberOfReps: sql.NullInt32{
			Int32: req.NumberOfReps,
			Valid: true,
		},
		NumberOfSets: sql.NullInt32{
			Int32: req.NumberOfSets,
			Valid: true,
		},
		RepetitionUnit: sql.NullString{String: req.RepetitionUnit, Valid: true},
		PhotoUrl:       sql.NullString{String: req.PhotoURL, Valid: true},
		VideoUrl:       sql.NullString{String: req.VideoURL, Valid: true},
	}

	exercise, err := c.store.CreateExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

type getExerciseRequest struct {
	ID string `uri:"id" binding:"required,min=1,uuid"`
}

func (c *exerciseController) GetExerciseByID(ctx *gin.Context) {
	var req getExerciseRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := c.store.GetExerciseById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

type updateExerciseRequest struct {
	Name           string `json:"Name" binding:"required"`
	Category       string `json:"category" binding:"required"`
	NumberOfReps   int32  `json:"NumberOfReps"`
	NumberOfSets   int32  `json:"NumberOfSets"`
	RepetitionUnit string `json:"RepetitionUnit"`
	PhotoURL       string `json:"PhotoURL"`
	VideoURL       string `json:"VideoURL"`
}

func (c *exerciseController) UpdateExerciseByID(ctx *gin.Context) {
	var req updateExerciseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := services.UpdateExerciseByIdParams{
		ID:       id,
		Name:     req.Name,
		Category: services.ExerciseCategory(req.Category),
		NumberOfReps: sql.NullInt32{
			Int32: req.NumberOfReps,
			Valid: true,
		},
		NumberOfSets: sql.NullInt32{
			Int32: req.NumberOfSets,
			Valid: true,
		},
		RepetitionUnit: sql.NullString{String: req.RepetitionUnit, Valid: true},
		PhotoUrl:       sql.NullString{String: req.PhotoURL, Valid: true},
		VideoUrl:       sql.NullString{String: req.VideoURL, Valid: true},
	}

	exercise, err := c.store.UpdateExerciseById(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (c *exerciseController) GetExercises(ctx *gin.Context) {
	exercises, err := c.store.GetExercises(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercises)
}
