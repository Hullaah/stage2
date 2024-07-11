-- Adds a user to a particular organisation
-- name: AddUserToOrganisation :exec
INSERT INTO membership (user_id, org_id)
VALUES ($1, $2);