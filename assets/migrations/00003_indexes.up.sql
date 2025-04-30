CREATE INDEX idx_user_id_active ON "user" (id)
WHERE
    deleted_at IS NULL;

CREATE INDEX idx_session_expires_at ON "session" (expires_at);

CREATE INDEX idx_image_user_active ON "image" (user_id)
WHERE
    deleted_at IS NULL;

CREATE INDEX idx_output_image_active ON "output" (image_id)
WHERE
    deleted_at IS NULL;

CREATE UNIQUE INDEX unique_output_path ON "output" (path)
WHERE
    deleted_at IS NULL;
