-- creates a new user and adds her to the user table
-- name: CreateUser :one
INSERT INTO "user" (first_name, last_name, email, password, phone)
VALUES ($1, $2, $3, $4, $5) RETURNING *;