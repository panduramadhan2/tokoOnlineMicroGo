CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone varchar(17) NULL,
    photo VARCHAR(255) NULL,
    address TEXT NULL,
    lat VARCHAR(50) NULL,
    lng VARCHAR(50) NULL,
    is_verified boolean DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- uploaded_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_is_verified ON users (is_verified);