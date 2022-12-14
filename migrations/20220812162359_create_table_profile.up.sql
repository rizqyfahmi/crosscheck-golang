CREATE TABLE IF NOT EXISTS "profiles" (
    id VARCHAR(255) UNIQUE PRIMARY KEY REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    photo TEXT NULL,
    date_of_birth DATE NULL,
    address TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL 
)