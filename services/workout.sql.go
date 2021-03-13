// Code generated by sqlc. DO NOT EDIT.
// source: workout.sql

package services

import (
	"context"

	"github.com/google/uuid"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO workout (
    name,
    split
) VALUES (
  $1, $2
) RETURNING id, created_at, updated_at, name, split
`

type CreateWorkoutParams struct {
	Name  string       `json:"name"`
	Split WorkoutSplit `json:"split"`
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout, arg.Name, arg.Split)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Split,
	)
	return i, err
}

const getWorkout = `-- name: GetWorkout :many
SELECT id, created_at, updated_at, name, split FROM workout
`

func (q *Queries) GetWorkout(ctx context.Context) ([]Workout, error) {
	rows, err := q.db.QueryContext(ctx, getWorkout)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workout{}
	for rows.Next() {
		var i Workout
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Split,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorkoutById = `-- name: GetWorkoutById :one
SELECT id, created_at, updated_at, name, split FROM workout WHERE id = $1
`

func (q *Queries) GetWorkoutById(ctx context.Context, id uuid.UUID) (Workout, error) {
	row := q.db.QueryRowContext(ctx, getWorkoutById, id)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Split,
	)
	return i, err
}

const updateWorkoutById = `-- name: UpdateWorkoutById :one
UPDATE workout SET name = $1, split = $2 WHERE id = $3 RETURNING id, created_at, updated_at, name, split
`

type UpdateWorkoutByIdParams struct {
	Name  string       `json:"name"`
	Split WorkoutSplit `json:"split"`
	ID    uuid.UUID    `json:"id"`
}

func (q *Queries) UpdateWorkoutById(ctx context.Context, arg UpdateWorkoutByIdParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, updateWorkoutById, arg.Name, arg.Split, arg.ID)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Split,
	)
	return i, err
}