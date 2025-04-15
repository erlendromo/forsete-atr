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

CREATE TABLE IF NOT EXISTS projects (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users_projects (
    user_id uuid,
    project_id uuid,
    PRIMARY KEY (user_id, project_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS images (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    project_id uuid,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS outputs (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    image_id uuid,
    FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE SET NULL
);
