-- Inserts a new organisation into the table and updates the membership table to reflect this update
-- name: CreateOrganisation :one
INSERT INTO organisation (name, description)
VALUES ($1, $2)
RETURNING *;