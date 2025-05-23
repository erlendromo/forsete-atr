CREATE INDEX IF NOT EXISTS idx_user_id_active ON "user" (id)
WHERE
    deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email_active ON "user" (email)
WHERE
    deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_session_expires_at ON "session" (expires_at);

CREATE INDEX IF NOT EXISTS idx_image_user_active ON "image" (user_id)
WHERE
    deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_output_image_active ON "output" (image_id)
WHERE
    deleted_at IS NULL;
