-- name: GetSpinner :one
SELECT * FROM spinners
WHERE spinners.UserID = $1 LIMIT 1;

-- name: GetSpinnerTricks :many
SELECT spinners.Tricks 
FROM spinners
INNER JOIN tricks 
ON spinners.Tricks = tricks.name
WHERE spinners.UserID = $1; 