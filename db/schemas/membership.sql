CREATE TABLE IF NOT EXISTS membership (
    user_id uuid,
    org_id uuid,
    CONSTRAINT user_fk FOREIGN KEY(user_id) REFERENCES "user"(user_id),
    CONSTRAINT org_fk FOREIGN KEY(org_id) REFERENCES organisation(org_id),
    PRIMARY KEY(user_id, org_id)
);
