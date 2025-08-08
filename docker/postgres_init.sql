CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     first_name TEXT NOT NULL,
                                     last_name  TEXT NOT NULL,
                                     email      TEXT UNIQUE NOT NULL
);

INSERT INTO users (first_name, last_name, email) VALUES
                                                     ('Ada','Lovelace','ada@example.com'),
                                                     ('Alan','Turing','alan@example.com')
    ON CONFLICT DO NOTHING;
    