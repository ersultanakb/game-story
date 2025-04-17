CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       is_admin BOOLEAN DEFAULT FALSE
);
