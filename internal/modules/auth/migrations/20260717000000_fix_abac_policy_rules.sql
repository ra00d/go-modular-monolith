-- +goose Up
-- Align the seeded ABAC policies with the lib/auth RuleSet schema the engine
-- actually parses (roles[], ownership{}, environment.time_range{}, ip_ranges[]).
-- The original init seed used a shorthand (role, time, ip_range) that the
-- engine cannot unmarshal, so those policies never matched.

UPDATE "abac_policy"
SET "rules" = '{
    "subject": {"roles": ["user"]},
    "resource": {"ownership": {"required": true, "field": "owner_id", "subject_field": "id"}},
    "environment": {}
}'
WHERE "slug" = 'own_content_edit';

UPDATE "abac_policy"
SET "rules" = '{
    "subject": {"roles": ["admin", "super_admin"]},
    "environment": {
        "time_range": {"start": "18:00", "end": "09:00", "timezone": "UTC"},
        "ip_ranges": ["10.0.0.0/8", "172.16.0.0/12"]
    }
}'
WHERE "slug" = 'admin_office_hours';

-- +goose Down

UPDATE "abac_policy"
SET "rules" = '{
    "subject": {"role": "user"},
    "resource": {"author_id": "{{subject.id}}"},
    "environment": {"time": {"start": "00:00", "end": "23:59"}}
}'
WHERE "slug" = 'own_content_edit';

UPDATE "abac_policy"
SET "rules" = '{
    "subject": {"role": "admin"},
    "environment": {
        "time": {"start": "18:00", "end": "09:00"},
        "ip_range": ["10.0.0.0/8", "172.16.0.0/12"]
    }
}'
WHERE "slug" = 'admin_office_hours';
