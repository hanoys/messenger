CREATE TABLE IF NOT EXISTS chats(
    id SERIAL PRIMARY KEY,
    users_id integer[]
);

