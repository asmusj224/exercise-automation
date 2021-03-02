-- name: GetExerciseById :one
SELECT * from exercise where id = $1;

-- name: CreateExercise :one
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
) RETURNING *;

-- name: UpdateExerciseById :one
UPDATE exercise SET 
  name = $1, 
  category = $2, 
  number_of_reps = $3,
  number_of_sets = $4,
  photo_url = $5,
  video_url = $6,
  repetition_unit = $7
  WHERE id = $8 RETURNING *;

-- name: GetExercises :many
SELECT * FROM exercise;
