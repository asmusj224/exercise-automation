// Code generated by sqlc. DO NOT EDIT.
// source: exercise.sql

package services

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createExercise = `-- name: CreateExercise :one
INSERT INTO exercise (
    category,
    name,
    number_of_reps,
    number_of_sets,
    repetition_unit,
    photo_url,
    video_url
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, created_at, updated_at, category, name, number_of_reps, number_of_sets, repetition_unit, photo_url, video_url
`

type CreateExerciseParams struct {
	Category       ExerciseCategory `json:"category"`
	Name           string           `json:"name"`
	NumberOfReps   int32            `json:"number_of_reps"`
	NumberOfSets   int32            `json:"number_of_sets"`
	RepetitionUnit sql.NullString   `json:"repetition_unit"`
	PhotoUrl       sql.NullString   `json:"photo_url"`
	VideoUrl       sql.NullString   `json:"video_url"`
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, createExercise,
		arg.Category,
		arg.Name,
		arg.NumberOfReps,
		arg.NumberOfSets,
		arg.RepetitionUnit,
		arg.PhotoUrl,
		arg.VideoUrl,
	)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Category,
		&i.Name,
		&i.NumberOfReps,
		&i.NumberOfSets,
		&i.RepetitionUnit,
		&i.PhotoUrl,
		&i.VideoUrl,
	)
	return i, err
}

const getExerciseById = `-- name: GetExerciseById :one
SELECT id, created_at, updated_at, category, name, number_of_reps, number_of_sets, repetition_unit, photo_url, video_url from exercise where id = $1
`

func (q *Queries) GetExerciseById(ctx context.Context, id uuid.UUID) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, getExerciseById, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Category,
		&i.Name,
		&i.NumberOfReps,
		&i.NumberOfSets,
		&i.RepetitionUnit,
		&i.PhotoUrl,
		&i.VideoUrl,
	)
	return i, err
}

const getExercises = `-- name: GetExercises :many
SELECT id, created_at, updated_at, category, name, number_of_reps, number_of_sets, repetition_unit, photo_url, video_url FROM exercise
`

func (q *Queries) GetExercises(ctx context.Context) ([]Exercise, error) {
	rows, err := q.db.QueryContext(ctx, getExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Exercise{}
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Category,
			&i.Name,
			&i.NumberOfReps,
			&i.NumberOfSets,
			&i.RepetitionUnit,
			&i.PhotoUrl,
			&i.VideoUrl,
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

const updateExerciseById = `-- name: UpdateExerciseById :one
UPDATE exercise SET 
  name = $1, 
  category = $2, 
  number_of_reps = $3,
  number_of_sets = $4,
  photo_url = $5,
  video_url = $6,
  repetition_unit = $7
  WHERE id = $8 RETURNING id, created_at, updated_at, category, name, number_of_reps, number_of_sets, repetition_unit, photo_url, video_url
`

type UpdateExerciseByIdParams struct {
	Name           string           `json:"name"`
	Category       ExerciseCategory `json:"category"`
	NumberOfReps   int32            `json:"number_of_reps"`
	NumberOfSets   int32            `json:"number_of_sets"`
	PhotoUrl       sql.NullString   `json:"photo_url"`
	VideoUrl       sql.NullString   `json:"video_url"`
	RepetitionUnit sql.NullString   `json:"repetition_unit"`
	ID             uuid.UUID        `json:"id"`
}

func (q *Queries) UpdateExerciseById(ctx context.Context, arg UpdateExerciseByIdParams) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, updateExerciseById,
		arg.Name,
		arg.Category,
		arg.NumberOfReps,
		arg.NumberOfSets,
		arg.PhotoUrl,
		arg.VideoUrl,
		arg.RepetitionUnit,
		arg.ID,
	)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Category,
		&i.Name,
		&i.NumberOfReps,
		&i.NumberOfSets,
		&i.RepetitionUnit,
		&i.PhotoUrl,
		&i.VideoUrl,
	)
	return i, err
}
