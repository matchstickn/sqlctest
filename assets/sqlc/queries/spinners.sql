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
FROM spinners
RETURN *;

--name: CreateSpinners :one
INSERT INTO spinners
(Name, Email, Provider, Tricks, ExpiresAt, AccessToken, AccessTokenSecret, RefreshToken)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURN *;

-- name: UpdateSpinners :one
UPDATE spinnners
SET Name = $2, Email = $3, Provider = $4, Tricks = $5, ExpiresAt = $6, AccessToken = $7, AccessTokenSecret = $8, RefreshToken = $9
WHERE id = $1
RETURN *;

-- name: DeleteSpinners :one
DELETE FROM spinnners
WHERE id = $1
RETURN *;