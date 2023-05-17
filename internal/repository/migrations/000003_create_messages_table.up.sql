CREATE TABLE IF NOT EXISTS messages(
    id SERIAL PRIMARY KEY,
    sender_id INTEGER,
    recipient_id INTEGER,
    time TIMESTAMP,
    body varchar(1024)
);
