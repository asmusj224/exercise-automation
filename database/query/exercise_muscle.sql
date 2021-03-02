-- name: GetExerciseMuscleById :one
SELECT * from exercise_muscle where id = $1;

-- name: CreateExerciseMuscle :one
INSERT INTO exercise_Muscle (
    exercise_id,
    muscle_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateExerciseMuscleById :one
UPDATE exercise_muscle SET exercise_id = $1, muscle_id = $2 WHERE id = $3 RETURNING *;
