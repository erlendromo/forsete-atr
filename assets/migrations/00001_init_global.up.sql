CREATE TABLE IF NOT EXISTS model_types (id serial PRIMARY KEY, type varchar(255));

CREATE TABLE IF NOT EXISTS models (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    model_type_id integer,
    FOREIGN KEY (model_type_id) REFERENCES model_types (id) ON DELETE SET NULL
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
