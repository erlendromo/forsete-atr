CREATE TABLE IF NOT EXISTS roles (
    id serial PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    email varchar(255) UNIQUE,
    password varchar(255) NOT NULL,
    role_id integer DEFAULT 2,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    token uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now (),
    expires_at timestamptz NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_sessions_expires_at ON sessions (expires_at);

CREATE TABLE IF NOT EXISTS images (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    user_id uuid,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS outputs (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    confirmed boolean,
    image_id uuid,
    FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE SET NULL
);
