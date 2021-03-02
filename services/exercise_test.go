package services

import (
	"context"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateExercise(t *testing.T) {
	ctx := context.Background()

	arg := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	exercise, err := testQueries.CreateExercise(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, arg.Category, exercise.Category)
	require.Equal(t, arg.Name, exercise.Name)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
}

func TestGetExerciseById(t *testing.T) {
	ctx := context.Background()

	arg := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	e, err := testQueries.CreateExercise(ctx, arg)

	exercise, err := testQueries.GetExerciseById(ctx, e.ID)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, e.ID, exercise.ID)
	require.Equal(t, e.CreatedAt, exercise.CreatedAt)
	require.Equal(t, e.Category, exercise.Category)
	require.Equal(t, e.Name, exercise.Name)
	require.Equal(t, e.UpdatedAt, exercise.UpdatedAt)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
}

func TestGetExercise(t *testing.T) {
	ctx := context.Background()

	arg := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	_, err := testQueries.CreateExercise(ctx, arg)

	_, err = testQueries.CreateExercise(ctx, arg1)

	exercises, err := testQueries.GetExercises(ctx)

	require.NoError(t, err)
	require.NotEmpty(t, exercises)
	require.Equal(t, len(exercises), 2)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise;")
}
func TestUpdateExerciseById(t *testing.T) {
	ctx := context.Background()

	arg := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	e, err := testQueries.CreateExercise(ctx, arg)

	updateArg := UpdateExerciseByIdParams{
		Name:     "Update Name",
		ID:       e.ID,
		Category: ExerciseCategoryCardio,
	}

	exercise, err := testQueries.UpdateExerciseById(ctx, updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, e.ID, exercise.ID)
	require.Equal(t, e.CreatedAt, exercise.CreatedAt)
	require.Equal(t, updateArg.Category, exercise.Category)
	require.Equal(t, updateArg.Name, exercise.Name)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
}
