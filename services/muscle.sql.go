// Code generated by sqlc. DO NOT EDIT.
// source: muscle.sql

package services

import (
	"context"

	"github.com/google/uuid"
)

const createMuscle = `-- name: CreateMuscle :one
INSERT INTO muscle (
    name
) VALUES (
  $1
) RETURNING id, created_at, updated_at, name
`

func (q *Queries) CreateMuscle(ctx context.Context, name string) (Muscle, error) {
	row := q.db.QueryRowContext(ctx, createMuscle, name)
	var i Muscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getMuscleById = `-- name: GetMuscleById :one
SELECT id, created_at, updated_at, name FROM muscle WHERE id = $1
`

func (q *Queries) GetMuscleById(ctx context.Context, id uuid.UUID) (Muscle, error) {
	row := q.db.QueryRowContext(ctx, getMuscleById, id)
	var i Muscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getMuscles = `-- name: GetMuscles :many
SELECT id, created_at, updated_at, name FROM muscle
`

func (q *Queries) GetMuscles(ctx context.Context) ([]Muscle, error) {
	rows, err := q.db.QueryContext(ctx, getMuscles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Muscle{}
	for rows.Next() {
		var i Muscle
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
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

const updateMuscleById = `-- name: UpdateMuscleById :one
UPDATE muscle SET name = $1 WHERE id = $2 RETURNING id, created_at, updated_at, name
`

type UpdateMuscleByIdParams struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateMuscleById(ctx context.Context, arg UpdateMuscleByIdParams) (Muscle, error) {
	row := q.db.QueryRowContext(ctx, updateMuscleById, arg.Name, arg.ID)
	var i Muscle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
