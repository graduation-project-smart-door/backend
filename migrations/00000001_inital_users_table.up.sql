BEGIN;

CREATE TABLE users (
       id TEXT NOT NULL PRIMARY KEY,
       email TEXT NOT NULL UNIQUE,
       first_name TEXT NOT NULL,
       last_name TEXT NOT NULL,
       patronymic TEXT,
       role TEXT NOT NULL,
       phone TEXT NOT NULL,
       password TEXT NOT NULL,

       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;