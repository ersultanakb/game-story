CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       description TEXT,
                       price NUMERIC(10, 2) NOT NULL
);
