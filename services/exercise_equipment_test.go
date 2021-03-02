package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateExerciseEquipment(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	equipment, err := testQueries.CreateEquipment(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseEquipmentParams{
		ExerciseID:  exercise.ID,
		EquipmentID: equipment.ID,
	}
	exerciseEquipment, err := testQueries.CreateExerciseEquipment(ctx, arg2)
	require.NoError(t, err)
	require.NotEmpty(t, equipment)
	require.NotEmpty(t, exercise)
	require.Equal(t, exerciseEquipment.EquipmentID, equipment.ID)
	require.Equal(t, exerciseEquipment.ExerciseID, exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment WHERE id = $1", equipment.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_equipment WHERE id = $1", exerciseEquipment.ID)
}

func TestGetExerciseEquipmentById(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	equipment, err := testQueries.CreateEquipment(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseEquipmentParams{
		ExerciseID:  exercise.ID,
		EquipmentID: equipment.ID,
	}
	exerciseEquipment, err := testQueries.CreateExerciseEquipment(ctx, arg2)
	foundExerciseEquipment, err := testQueries.GetExerciseEquipmentById(ctx, exerciseEquipment.ID)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseEquipment.ID, foundExerciseEquipment.ID)
	require.Equal(t, exerciseEquipment.ExerciseID, foundExerciseEquipment.ExerciseID)
	require.Equal(t, exerciseEquipment.EquipmentID, foundExerciseEquipment.EquipmentID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment WHERE id = $1", equipment.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_equipment WHERE id = $1", exerciseEquipment.ID)
}

func TestUpdateExerciseEquipmentById(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	equipment, err := testQueries.CreateEquipment(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)

	eq2 := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}

	ex2 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	equipment2, err := testQueries.CreateEquipment(ctx, eq2)
	exercise2, err := testQueries.CreateExercise(ctx, ex2)
	arg2 := CreateExerciseEquipmentParams{
		ExerciseID:  exercise.ID,
		EquipmentID: equipment.ID,
	}
	exerciseEquipment, err := testQueries.CreateExerciseEquipment(ctx, arg2)
	updateParams := UpdateExerciseEquipmentByIdParams{
		ID:          exerciseEquipment.ID,
		ExerciseID:  exercise2.ID,
		EquipmentID: equipment2.ID,
	}
	updatedExerciseEquipment, err := testQueries.UpdateExerciseEquipmentById(ctx, updateParams)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseEquipment.ID, updatedExerciseEquipment.ID)
	require.Equal(t, exercise2.ID, updatedExerciseEquipment.ExerciseID)
	require.Equal(t, equipment2.ID, updatedExerciseEquipment.EquipmentID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_equipment WHERE id = $1", exerciseEquipment.ID)
}
