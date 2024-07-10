CREATE TABLE IF  NOT EXISTS organisation (
    org_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);
