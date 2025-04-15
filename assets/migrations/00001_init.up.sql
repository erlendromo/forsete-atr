CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    email varchar(255) UNIQUE,
    encrypted_password varchar(255) NOT NULL
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

CREATE TABLE IF NOT EXISTS model_type (id serial PRIMARY KEY, type varchar(255));

CREATE TABLE IF NOT EXISTS models (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    model_type_id integer,
    FOREIGN KEY (model_type_id) REFERENCES model_type (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS pipelines (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS pipelines_models (
    pipeline_id integer,
    model_id integer,
    PRIMARY KEY (pipeline_id, model_id),
    FOREIGN KEY (pipeline_id) REFERENCES pipelines (id) ON DELETE CASCADE,
    FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE CASCADE
);
