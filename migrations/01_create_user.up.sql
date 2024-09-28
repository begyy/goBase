CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    username     VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL,
    firstname    VARCHAR(255),
    lastname     VARCHAR(255),
    password     VARCHAR(255) NULL,
    is_superuser BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);
