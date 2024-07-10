-- Gets a user in same organisation with a particular user
SELECT user_id,
    first_name,
    last_name,
    email,
    phone
FROM user
WHERE id = "user id in path variable"
    AND EXISTS (
        SELECT 1
        FROM organisation o1
            INNER JOIN membership m1 USING(org_id)
        WHERE m1.user_id = "user id in path variable"
        INTERSECT
        SELECT 1
        FROM organisation o2
            INNER JOIN membership m2 USING(org_id)
        WHERE m2.user_id = "logged in user id"
    );