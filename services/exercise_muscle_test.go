package services

import (
	"context"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateExerciseMuscle(t *testing.T) {
	ctx := context.Background()

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	muscle, err := testQueries.CreateMuscle(ctx, util.RandomName())
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseMuscleParams{
		ExerciseID: exercise.ID,
		MuscleID:   muscle.ID,
	}
	exerciseMuscle, err := testQueries.CreateExerciseMuscle(ctx, arg2)
	require.NoError(t, err)
	require.NotEmpty(t, muscle)
	require.NotEmpty(t, exercise)
	require.Equal(t, exerciseMuscle.MuscleID, muscle.ID)
	require.Equal(t, exerciseMuscle.ExerciseID, exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM muscle WHERE id = $1", muscle.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_muscle WHERE id = $1", exerciseMuscle.ID)
}

func TestGetExerciseMuscleById(t *testing.T) {
	ctx := context.Background()
	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	muscle, err := testQueries.CreateMuscle(ctx, util.RandomName())
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseMuscleParams{
		ExerciseID: exercise.ID,
		MuscleID:   muscle.ID,
	}
	exerciseMuscle, err := testQueries.CreateExerciseMuscle(ctx, arg2)
	foundExerciseMuscle, err := testQueries.GetExerciseMuscleById(ctx, exerciseMuscle.ID)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseMuscle.ID, foundExerciseMuscle.ID)
	require.Equal(t, exerciseMuscle.ExerciseID, foundExerciseMuscle.ExerciseID)
	require.Equal(t, exerciseMuscle.MuscleID, foundExerciseMuscle.MuscleID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM muscle WHERE id = $1", muscle.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_muscle WHERE id = $1", exerciseMuscle.ID)
}

func TestUpdateExerciseMuscleById(t *testing.T) {
	ctx := context.Background()

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	muscle, err := testQueries.CreateMuscle(ctx, util.RandomName())
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	ex2 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	muscle2, err := testQueries.CreateMuscle(ctx, util.RandomName())
	exercise2, err := testQueries.CreateExercise(ctx, ex2)
	arg2 := CreateExerciseMuscleParams{
		ExerciseID: exercise.ID,
		MuscleID:   muscle.ID,
	}
	exerciseMuscle, err := testQueries.CreateExerciseMuscle(ctx, arg2)
	updateParams := UpdateExerciseMuscleByIdParams{
		ID:         exerciseMuscle.ID,
		ExerciseID: exercise2.ID,
		MuscleID:   muscle2.ID,
	}
	updatedExerciseMuscle, err := testQueries.UpdateExerciseMuscleById(ctx, updateParams)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseMuscle.ID, updatedExerciseMuscle.ID)
	require.Equal(t, exercise2.ID, updatedExerciseMuscle.ExerciseID)
	require.Equal(t, muscle2.ID, updatedExerciseMuscle.MuscleID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM muscle")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_muscle WHERE id = $1", exerciseMuscle.ID)
}
