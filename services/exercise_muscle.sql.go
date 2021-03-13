// Code generated by sqlc. DO NOT EDIT.
// source: exercise_muscle.sql

package services

import (
	"context"

	"github.com/google/uuid"
)

const createExerciseMuscle = `-- name: CreateExerciseMuscle :one
INSERT INTO exercise_Muscle (
    exercise_id,
    muscle_id
) VALUES (
  $1, $2
) RETURNING id, created_at, updated_at, exercise_id, muscle_id
`

type CreateExerciseMuscleParams struct {
	ExerciseID uuid.UUID `json:"exercise_id"`
	MuscleID   uuid.UUID `json:"muscle_id"`
}

func (q *Queries) CreateExerciseMuscle(ctx context.Context, arg CreateExerciseMuscleParams) (ExerciseMuscle, error) {
	row := q.db.QueryRowContext(ctx, createExerciseMuscle, arg.ExerciseID, arg.MuscleID)
	var i ExerciseMuscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExerciseID,
		&i.MuscleID,
	)
	return i, err
}

const getExerciseMuscleById = `-- name: GetExerciseMuscleById :one
SELECT id, created_at, updated_at, exercise_id, muscle_id from exercise_muscle where id = $1
`

func (q *Queries) GetExerciseMuscleById(ctx context.Context, id uuid.UUID) (ExerciseMuscle, error) {
	row := q.db.QueryRowContext(ctx, getExerciseMuscleById, id)
	var i ExerciseMuscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExerciseID,
		&i.MuscleID,
	)
	return i, err
}

const updateExerciseMuscleById = `-- name: UpdateExerciseMuscleById :one
UPDATE exercise_muscle SET exercise_id = $1, muscle_id = $2 WHERE id = $3 RETURNING id, created_at, updated_at, exercise_id, muscle_id
`

type UpdateExerciseMuscleByIdParams struct {
	ExerciseID uuid.UUID `json:"exercise_id"`
	MuscleID   uuid.UUID `json:"muscle_id"`
	ID         uuid.UUID `json:"id"`
}

func (q *Queries) UpdateExerciseMuscleById(ctx context.Context, arg UpdateExerciseMuscleByIdParams) (ExerciseMuscle, error) {
	row := q.db.QueryRowContext(ctx, updateExerciseMuscleById, arg.ExerciseID, arg.MuscleID, arg.ID)
	var i ExerciseMuscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExerciseID,
		&i.MuscleID,
	)
	return i, err
}