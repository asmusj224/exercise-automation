-- name: GetExerciseEquipmentById :one
SELECT * from exercise_equipment where id = $1;

-- name: CreateExerciseEquipment :one
INSERT INTO exercise_equipment (
    exercise_id,
    equipment_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateExerciseEquipmentById :one
UPDATE exercise_equipment SET exercise_id = $1, equipment_id = $2 WHERE id = $3 RETURNING *;

