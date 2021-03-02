-- name: CreateEquipment :one
INSERT INTO equipment (
    name
) VALUES (
  $1
) RETURNING *;


-- name: GetEquipmentById :one
SELECT * FROM equipment where id = $1;

-- name: UpdateEquipmentById :one
UPDATE equipment SET name = $1 WHERE id = $2 RETURNING *;

-- name: GetEquipment :many
SELECT * FROM equipment;
