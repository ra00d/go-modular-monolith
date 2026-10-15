-- +goose Up
ALTER TABLE abac_policy
   ALTER COLUMN rules TYPE JSONB USING rules::jsonb;

-- +goose Down
ALTER TABLE abac_policy
   ALTER COLUMN rules TYPE TEXT USING rules::text;


