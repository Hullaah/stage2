-- Creates an organisation and adds it to the organisation table
-- name: CreateOrganisation :one
INSERT INTO organisation (name, description)
VALUES ($1, $2)
RETURNING *;