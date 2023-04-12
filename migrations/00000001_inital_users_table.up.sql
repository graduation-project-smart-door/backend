BEGIN;

CREATE TYPE roles AS ENUM ('admin', 'superadmin', 'user');

CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       person_id TEXT NOT NULL,
       email TEXT UNIQUE,
       first_name TEXT NOT NULL,
       patronymic TEXT,
       last_name TEXT NOT NULL,
       role roles NOT NULL,
       phone TEXT,
       password TEXT,
       position TEXT NOT NULL,
       avatar TEXT,

       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;