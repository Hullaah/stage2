-- Gets a specific organisation which a user belongs to 
-- name: GetUserOrganisation :one
SELECT org.org_id,
    org.name,
    org.description
FROM organisation org
    INNER JOIN membership m ON org.org_id = m.org_id
WHERE m.user_id = "logged in user id"
    AND org.org_id = "org id from path variable";