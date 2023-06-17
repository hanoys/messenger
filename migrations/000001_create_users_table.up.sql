CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    nickname TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP
);

