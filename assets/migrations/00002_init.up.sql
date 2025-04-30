CREATE TABLE IF NOT EXISTS "role" (
    id serial PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamptz DEFAULT now (),
    deleted_at timestamptz,
    role_id integer NOT NULL DEFAULT 2,
    FOREIGN KEY (role_id) REFERENCES "role" (id)
);

CREATE TABLE IF NOT EXISTS "session" (
    token uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid NOT NULL,
    created_at timestamptz DEFAULT now (),
    expires_at timestamptz NOT NULL DEFAULT (now () + interval '24 hours'),
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE INDEX idx_session_expires_at ON "session" (expires_at);

CREATE TABLE IF NOT EXISTS "image" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) UNIQUE NOT NULL,
    format varchar(10) NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    uploaded_at timestamptz DEFAULT now (),
    deleted_at timestamptz,
    user_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE IF NOT EXISTS "output" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) UNIQUE NOT NULL,
    format varchar(10) NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    created_at timestamptz DEFAULT now (),
    updated_at timestamptz DEFAULT now (),
    deleted_at timestamptz,
    confirmed bool,
    image_id uuid NOT NULL,
    FOREIGN KEY (image_id) REFERENCES "image" (id)
);
