package services

import (
	"context"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateMuscle(t *testing.T) {
	ctx := context.Background()
	name := util.RandomName()

	muscle, err := testQueries.CreateMuscle(ctx, name)
	require.NoError(t, err)
	require.NotEmpty(t, muscle)
	require.Equal(t, name, muscle.Name)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM muscle WHERE id = $1", muscle.ID)
}

func TestGetMuscleById(t *testing.T) {
	ctx := context.Background()
	name := util.RandomName()

	e, err := testQueries.CreateMuscle(ctx, name)
	muscle, err := testQueries.GetMuscleById(ctx, e.ID)
	require.NoError(t, err)
	require.NotEmpty(t, muscle)
	require.Equal(t, e.ID, muscle.ID)
	require.Equal(t, e.CreatedAt, muscle.CreatedAt)
	require.Equal(t, e.Name, muscle.Name)
	require.Equal(t, e.UpdatedAt, muscle.UpdatedAt)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM muscle WHERE id = $1", muscle.ID)
}

func TestGetMuscle(t *testing.T) {
	ctx := context.Background()

	_, err := testQueries.CreateMuscle(ctx, util.RandomName())

	_, err = testQueries.CreateMuscle(ctx, util.RandomName())

	Muscles, err := testQueries.GetMuscles(ctx)

	require.NoError(t, err)
	require.NotEmpty(t, Muscles)
	require.Equal(t, len(Muscles), 2)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM Muscle;")
}
func TestUpdateMuscleById(t *testing.T) {
	ctx := context.Background()

	name := util.RandomName()

	e, err := testQueries.CreateMuscle(ctx, name)

	updateArg := UpdateMuscleByIdParams{
		Name: "Update Name",
		ID:   e.ID,
	}

	muscle, err := testQueries.UpdateMuscleById(ctx, updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, muscle)
	require.Equal(t, e.ID, muscle.ID)
	require.Equal(t, e.CreatedAt, muscle.CreatedAt)
	require.Equal(t, updateArg.Name, muscle.Name)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM Muscle WHERE id = $1", muscle.ID)
}
