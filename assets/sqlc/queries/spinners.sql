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
(Name, Email, AdminPerms, Provider, Tricks, ExpiresAt, AccessToken, AccessTokenSecret, RefreshToken)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateSpinner :one
UPDATE spinners
SET Name = $2, Email = $3, AdminPerms = $4, Provider = $5, Tricks = $6, ExpiresAt = $7, AccessToken = $8, AccessTokenSecret = $9, RefreshToken = $10
WHERE UserID = $1
RETURNING *;

-- name: DeleteSpinner :one
DELETE FROM spinners
WHERE UserID = $1
RETURNING *;

-- name: RetriveSpinner :one
SELECT * FROM spinners
WHERE Accesstoken = $1 LIMIT 1;