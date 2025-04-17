CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER REFERENCES users(id),
                       title TEXT NOT NULL,
                       description TEXT,
                       price NUMERIC(10, 2) NOT NULL
);
