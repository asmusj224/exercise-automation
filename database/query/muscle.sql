-- name: CreateMuscle :one
INSERT INTO muscle (
    name
) VALUES (
  $1
) RETURNING *;

-- name: GetMuscleById :one
SELECT * FROM muscle WHERE id = $1;

-- name: UpdateMuscleById :one
UPDATE muscle SET name = $1 WHERE id = $2 RETURNING *;

-- name: GetMuscles :many
SELECT * FROM muscle;
