package services

import (
	"context"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkout(t *testing.T) {
	arg := CreateWorkoutParams{
		Name:  "Test Workout",
		Split: WorkoutSplitFullBody,
	}
	ctx := context.Background()

	workout, err := testQueries.CreateWorkout(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.Equal(t, arg.Name, workout.Name)
	require.Equal(t, arg.Split, workout.Split)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout WHERE id = $1", workout.ID)
}

func TestGetWorkoutById(t *testing.T) {
	arg := CreateWorkoutParams{
		Name:  "Test Workout",
		Split: WorkoutSplitFullBody,
	}
	ctx := context.Background()

	w, err := testQueries.CreateWorkout(ctx, arg)

	workout, err := testQueries.GetWorkoutById(ctx, w.ID)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.Equal(t, w.ID, workout.ID)
	require.Equal(t, arg.Name, workout.Name)
	require.Equal(t, arg.Split, workout.Split)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout WHERE id = $1", workout.ID)
}

func TestGetWorkout(t *testing.T) {
	arg := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}

	arg2 := CreateWorkoutParams{
		Name:  util.RandomName(),
		Split: WorkoutSplitFullBody,
	}
	ctx := context.Background()

	_, err := testQueries.CreateWorkout(ctx, arg)
	_, err = testQueries.CreateWorkout(ctx, arg2)

	workouts, err := testQueries.GetWorkout(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, workouts)
	require.Equal(t, len(workouts), 2)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout;")
}

func TestUpdateWorkoutById(t *testing.T) {
	arg := CreateWorkoutParams{
		Name:  "Test Workout",
		Split: WorkoutSplitFullBody,
	}
	ctx := context.Background()

	w, err := testQueries.CreateWorkout(ctx, arg)

	updateArg := UpdateWorkoutByIdParams{
		Name:  "Update Name",
		Split: arg.Split,
		ID:    w.ID,
	}

	workout, err := testQueries.UpdateWorkoutById(ctx, updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.Equal(t, w.ID, workout.ID)
	require.Equal(t, updateArg.Name, workout.Name)
	require.Equal(t, arg.Split, workout.Split)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM workout WHERE id = $1", workout.ID)
}
