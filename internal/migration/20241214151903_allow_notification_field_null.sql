-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.notifications
ALTER COLUMN reference_id DROP NOT NULL;

ALTER TABLE public.notifications
ALTER COLUMN reference_type DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.notifications
ALTER COLUMN reference_id SET NOT NULL;

ALTER TABLE public.notifications
ALTER COLUMN reference_type SET NOT NULL;
-- +goose StatementEnd