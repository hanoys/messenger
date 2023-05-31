CREATE TABLE IF NOT EXISTS messages(
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL REFERENCES users,
    recipient_id INTEGER NOT NULL REFERENCES users,
    chat_id INTEGER NOT NULL REFERENCES chats,
    time TIMESTAMP,
    body varchar(1024)
);
