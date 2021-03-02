package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/asmusj224/exercise-automation/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEquipment(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}
	equipment, err := testQueries.CreateEquipment(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, equipment)
	require.Equal(t, arg.String, equipment.Name.String)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment WHERE id = $1", equipment.ID)
}

func TestGetEquipment(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}
	arg1 := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}
	_, err := testQueries.CreateEquipment(ctx, arg)
	require.NoError(t, err)
	_, err = testQueries.CreateEquipment(ctx, arg1)
	require.NoError(t, err)
	equipment, err := testQueries.GetEquipment(ctx)
	require.NoError(t, err)
	require.Equal(t, len(equipment), 2)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment")
}

func TestGetEquipmentById(t *testing.T) {
	ctx := context.Background()
	arg := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}
	e, err := testQueries.CreateEquipment(ctx, arg)
	require.NoError(t, err)
	equipment, err := testQueries.GetEquipmentById(ctx, e.ID)
	require.NoError(t, err)
	require.Equal(t, e.ID, equipment.ID)
	require.Equal(t, e.Name.String, equipment.Name.String)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment")
}

func TestUpdateEquipmentById(t *testing.T) {
	ctx := context.Background()
	name := sql.NullString{
		String: util.RandomName(),
		Valid:  true,
	}
	e, err := testQueries.CreateEquipment(ctx, name)
	require.NoError(t, err)
	arg := UpdateEquipmentByIdParams{
		Name: sql.NullString{
			String: "Update",
			Valid:  true,
		},
		ID: e.ID,
	}
	equipment, err := testQueries.UpdateEquipmentById(ctx, arg)
	require.NoError(t, err)
	require.Equal(t, e.ID, equipment.ID)
	require.Equal(t, arg.Name.String, equipment.Name.String)
	testQueries.db.QueryRowContext(ctx, "DELETE FROM equipment")
}
