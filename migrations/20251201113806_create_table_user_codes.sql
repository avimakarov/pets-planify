-- +goose Up
-- +goose StatementBegin
create table if not exists user_codes (
    id uuid not null default gen_random_uuid() primary key,
    code uuid not null,
    user_id uuid not null references users(id),
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_code;
-- +goose StatementEnd
