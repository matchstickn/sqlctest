-- name: GetTrick :one
SELECT * FROM tricks
WHERE id = $1 LIMIT 1;

-- name: GetAllTricks :many
SELECT * FROM tricks 
ORDER BY name;

-- name: CreateTrick :one
INSERT INTO tricks 
(name, style, power) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: UpdateTrick :one
UPDATE tricks
SET name = $2, style = $3, power = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTrick :exec
DELETE FROM tricks
WHERE id = $1;