CREATE TABLE IF NOT EXISTS "user" (
    user_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    phone CHAR(11)
);