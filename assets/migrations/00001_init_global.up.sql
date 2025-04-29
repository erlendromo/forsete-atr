CREATE TABLE IF NOT EXISTS "model_type" (
    id serial PRIMARY KEY,
    type varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "model" (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL,
    model_type_id integer NOT NULL,
    FOREIGN KEY (model_type_id) REFERENCES "model_type" (id)
);

CREATE TABLE IF NOT EXISTS "pipeline" (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    path varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "pipeline_model" (
    pipeline_id integer,
    model_id integer,
    PRIMARY KEY (pipeline_id, model_id),
    FOREIGN KEY (pipeline_id) REFERENCES "pipeline" (id),
    FOREIGN KEY (model_id) REFERENCES "model" (id)
);
