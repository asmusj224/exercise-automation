-- name: GetExerciseWorkoutById :one
SELECT * from exercise_workout where id = $1;

-- name: CreateExerciseWorkout :one
INSERT INTO exercise_workout (
    exercise_id,
    workout_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateExerciseWorkoutById :one
UPDATE exercise_workout SET exercise_id = $1, workout_id = $2 WHERE id = $3 RETURNING *;


-- name: GetExerciseWorkoutByWorkoutId :one
SELECT w.id, w.name, w.split , t.exercises
FROM   workout      w
JOIN  ( 
   SELECT ew.workout_id AS id, array_to_json(array_agg(e.*)) AS exercises
   FROM   exercise_workout ew
   JOIN   exercise       e  ON e.id = ew.exercise_id
   GROUP  BY ew.workout_id
   ) t USING (id) WHERE w.id = $1;