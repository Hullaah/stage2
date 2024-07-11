-- Gets a user
-- name: GetUser :one
SELECT *
FROM "user"
WHERE email = $1;