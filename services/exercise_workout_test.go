package services

import (
	"context"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateExerciseWorkout(t *testing.T) {
	ctx := context.Background()

	arg := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	workout, err := testQueries.CreateWorkout(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseWorkoutParams{
		ExerciseID: exercise.ID,
		WorkoutID:  workout.ID,
	}
	exerciseWorkout, err := testQueries.CreateExerciseWorkout(ctx, arg2)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.NotEmpty(t, exercise)
	require.Equal(t, exerciseWorkout.WorkoutID, workout.ID)
	require.Equal(t, exerciseWorkout.ExerciseID, exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout WHERE id = $1", workout.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_Workout WHERE id = $1", exerciseWorkout.ID)
}

func TestGetExerciseWorkoutById(t *testing.T) {
	ctx := context.Background()

	arg := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	workout, err := testQueries.CreateWorkout(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	arg2 := CreateExerciseWorkoutParams{
		ExerciseID: exercise.ID,
		WorkoutID:  workout.ID,
	}
	exerciseWorkout, err := testQueries.CreateExerciseWorkout(ctx, arg2)
	foundExerciseWorkout, err := testQueries.GetExerciseWorkoutById(ctx, exerciseWorkout.ID)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseWorkout.ID, foundExerciseWorkout.ID)
	require.Equal(t, exerciseWorkout.ExerciseID, foundExerciseWorkout.ExerciseID)
	require.Equal(t, exerciseWorkout.WorkoutID, foundExerciseWorkout.WorkoutID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM Workout WHERE id = $1", workout.ID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_workout WHERE id = $1", exerciseWorkout.ID)
}

func TestUpdateExerciseWorkoutById(t *testing.T) {
	ctx := context.Background()

	arg := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}

	arg1 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}

	workout, err := testQueries.CreateWorkout(ctx, arg)
	exercise, err := testQueries.CreateExercise(ctx, arg1)
	ex2 := CreateExerciseParams{
		Category: ExerciseCategoryHIIT,
		Name:     util.RandomName(),
	}
	wk := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}

	workout2, err := testQueries.CreateWorkout(ctx, wk)
	exercise2, err := testQueries.CreateExercise(ctx, ex2)
	arg2 := CreateExerciseWorkoutParams{
		ExerciseID: exercise.ID,
		WorkoutID:  workout.ID,
	}
	exerciseWorkout, err := testQueries.CreateExerciseWorkout(ctx, arg2)
	updateParams := UpdateExerciseWorkoutByIdParams{
		ID:         exerciseWorkout.ID,
		ExerciseID: exercise2.ID,
		WorkoutID:  workout2.ID,
	}
	updatedExerciseWorkout, err := testQueries.UpdateExerciseWorkoutById(ctx, updateParams)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, exerciseWorkout.ID, updatedExerciseWorkout.ID)
	require.Equal(t, exercise2.ID, updatedExerciseWorkout.ExerciseID)
	require.Equal(t, workout2.ID, updatedExerciseWorkout.WorkoutID)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout")
	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_workout WHERE id = $1", exerciseWorkout.ID)
}

// func TestGetExerciseWorkoutByWorkoutId(t *testing.T) {
// 	ctx := context.Background()

// 	arg := CreateWorkoutParams{
// 		Name:  util.RandomName(),
// 		Split: WorkoutSplitFullBody,
// 	}

// 	arg1 := CreateExerciseParams{
// 		Category: ExerciseCategoryHIIT,
// 		Name:     util.RandomName(),
// 	}

// 	workout, err := testQueries.CreateWorkout(ctx, arg)
// 	exercise, err := testQueries.CreateExercise(ctx, arg1)
// 	arg2 := CreateExerciseWorkoutParams{
// 		ExerciseID: exercise.ID,
// 		WorkoutID:  workout.ID,
// 	}
// 	exerciseWorkout, err := testQueries.CreateExerciseWorkout(ctx, arg2)
// 	foundExerciseWorkout, err := testQueries.GetExerciseWorkoutByWorkoutId(ctx, exerciseWorkout.ID)
// 	require.NoError(t, err)
// 	require.NoError(t, err)
// 	require.Empty(t, foundExerciseWorkout)
// 	// require.Equal(t, exerciseWorkout.ID, foundExerciseWorkout.ID)
// 	// require.Equal(t, exerciseWorkout.ExerciseID, foundExerciseWorkout.ExerciseID)
// 	// require.Equal(t, exerciseWorkout.WorkoutID, foundExerciseWorkout.WorkoutID)
// 	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise WHERE id = $1", exercise.ID)
// 	testQueries.db.QueryRowContext(ctx, "DELETE FROM Workout WHERE id = $1", workout.ID)
// 	testQueries.db.QueryRowContext(ctx, "DELETE FROM exercise_workout WHERE id = $1", exerciseWorkout.ID)
// }
