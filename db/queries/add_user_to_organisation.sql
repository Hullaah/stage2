-- name: AddUserToOrganisation :exec
INSERT INTO membership (user_id, org_id)
VALUES (
        "user id from request bosy",
        "organisation id from path variable"
    );