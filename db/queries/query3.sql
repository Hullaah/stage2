-- Gets all the organisations a user belongs to
-- name: GetUserOrganisations :many
SELECT org.org_id,
    org.name,
    org.description
FROM organisation org
    INNER JOIN membership m ON org.org_id = m.org_id
WHERE m.user_id = "logged in user id";