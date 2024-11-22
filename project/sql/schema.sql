CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

INSERT INTO users (name, email) VALUES
('John Doe', 'john@example.com'),
('Jane Doe', 'jane@example.com');
