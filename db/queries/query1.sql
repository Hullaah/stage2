-- Gets a user in same organisation with a particular user
-- name: GetUserIfInSameOrganisation :one
SELECT u.user_id,
    u.first_name,
    u.last_name,
    u.email,
    u.phone
FROM "user" u
    JOIN membership m1 ON u.user_id = m1.user_id
    JOIN membership m2 ON m1.org_id = m2.org_id
WHERE u.user_id = $1
    AND m2.user_id = $2;