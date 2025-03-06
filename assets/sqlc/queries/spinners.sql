-- name: GetSpinner :one
SELECT * FROM spinners
WHERE spinners.UserID = $1 LIMIT 1;

-- name: GetSpinnerTricks :many
SELECT spinners.Tricks 
FROM spinners
INNER JOIN tricks 
ON spinners.Tricks = tricks.name
WHERE spinners.UserID = $1; 

-- name: ListSpinners :many
SELECT *
FROM spinners;

-- name: CreateSpinner :one
INSERT INTO spinners
(Name, Email, Provider, Tricks, ExpiresAt, AccessToken, AccessTokenSecret, RefreshToken)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateSpinner :one
UPDATE spinners
SET Name = $2, Email = $3, Provider = $4, Tricks = $5, ExpiresAt = $6, AccessToken = $7, AccessTokenSecret = $8, RefreshToken = $9
WHERE UserID = $1
RETURNING *;

-- name: DeleteSpinner :one
DELETE FROM spinners
WHERE UserID = $1
RETURNING *;