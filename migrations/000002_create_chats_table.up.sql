DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'chat_type') THEN
        CREATE TYPE chat_type AS ENUM ('conversation', 'group');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS chats(
    id SERIAL PRIMARY KEY,
    name TEXT,
    type chat_type
);

CREATE TABLE IF NOT EXISTS chats_users(
    user_id INTEGER REFERENCES users,
    chat_id INTEGER REFERENCES chats
)

