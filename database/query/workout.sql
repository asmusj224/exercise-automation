-- name: CreateWorkout :one
INSERT INTO workout (
    name,
    split
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetWorkoutById :one
SELECT * FROM workout WHERE id = $1;

-- name: UpdateWorkoutById :one
UPDATE workout SET name = $1, split = $2 WHERE id = $3 RETURNING *;

-- name: GetWorkout :many
SELECT * FROM workout;
