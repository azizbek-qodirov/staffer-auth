CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    position_id UUID DEFAULT NULL,
    rating FLOAT,
    role VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE refresh_tokens (
    user_email,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP,
);
