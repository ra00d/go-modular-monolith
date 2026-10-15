-- +goose Up
ALTER TABLE "user" ADD COLUMN preferences JSONB NOT NULL DEFAULT '{}'::jsonb;
-- +goose Down
ALTER TABLE "user" DROP COLUMN preferences;


