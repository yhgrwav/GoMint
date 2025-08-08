CREATE TABLE IF NOT EXISTS users (
                                     id BIGINT PRIMARY KEY AUTO_INCREMENT,
                                     first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE
    );

INSERT INTO users (first_name, last_name, email) VALUES
                                                     ('Ada','Lovelace','ada@example.com'),
                                                     ('Alan','Turing','alan@example.com')
    ON DUPLICATE KEY UPDATE email = VALUES(email);
