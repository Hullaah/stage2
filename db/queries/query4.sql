-- Inserts a new organisation into the table and updates the membership table to reflect this update
-- 
INSERT INTO organisation (name, description)
VALUES (
        "name from request body",
        "description from request body"
    )
RETURNING org_id;
INSERT INTO membership (org_id, user_id)
VALUES (
        "organisation id",
        "logged in user id"
    );