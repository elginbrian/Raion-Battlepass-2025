package migrations

const CreatePostsTable = `
CREATE EXTENSION IF NOT EXISTS pgcrypto; 

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,  
    image_url TEXT NOT NULL,
    caption TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, image_url, caption) 
);
`

const InsertPostsData = `
INSERT INTO posts (user_id, image_url, caption, created_at, updated_at)
VALUES 
    ((SELECT id FROM users WHERE email = 'elginbrian49@gmail.com'), 'https://raion-battlepass.elginbrian.com/uploads/hackjam.jpg', 'With fellas at Raion Hackjam 2024.', NOW(), NOW())
ON CONFLICT (user_id, image_url, caption) DO NOTHING;

INSERT INTO posts (user_id, image_url, caption, created_at, updated_at)
VALUES 
    ((SELECT id FROM users WHERE email = 'midnightsparks@example.com'), 'https://raion-battlepass.elginbrian.com/uploads/farewell.jpg', 'Our activities at Raion Farewell 2024.', NOW(), NOW())
ON CONFLICT (user_id, image_url, caption) DO NOTHING;
`
