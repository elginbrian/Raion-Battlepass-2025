package migrations

const CreateUsersTable = `
CREATE EXTENSION IF NOT EXISTS pgcrypto; 

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

`
const InsertUsersData = `
INSERT INTO users (name, email, password_hash, created_at, updated_at)
VALUES 
    ('elginbrian', 'elginbrian49@gmail.com', '$2a$10$D1g0OZLqH1rO5Gp2f9D5Fq2tqJb1h0B0VJhgV9AqE6qL6q8XgXy8G', NOW(), NOW())
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (name, email, password_hash, created_at, updated_at)
VALUES 
    ('midnightsparks', 'midnightsparks@example.com', '$2a$10$E1g0OZLqH1rO5Gp2f9D5Fq2tqJb1h0B0VJhgV9AqE6qL6q8XgXy8G', NOW(), NOW())
ON CONFLICT (email) DO NOTHING;
`