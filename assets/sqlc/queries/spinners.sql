-- name: GetUser :one
SELECT * FROM spinners
WHERE spinners.UserID = $1 LIMIT 1;

-- name: GetUserTricks :many
SELECT spinners.Tricks 
FROM spinners
INNER JOIN tricks 
ON spinners.Tricks = tricks.name;