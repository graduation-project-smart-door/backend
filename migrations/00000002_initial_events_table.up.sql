BEGIN;

CREATE TYPE event_types AS ENUM ('exit', 'enter');

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    direction event_types NOT NULL,
    user_id INT NOT NULL,
    event_time TIMESTAMPTZ,

    CONSTRAINT fk_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
);

COMMIT;