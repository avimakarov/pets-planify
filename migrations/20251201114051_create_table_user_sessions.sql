-- +goose Up
-- +goose StatementBegin
create table if not exists user_sessions (
    id uuid not null default gen_random_uuid() primary key,
    token uuid not null unique,
    user_id uuid not null references users(id),
    code_id uuid not null references user_codes(id),
    is_active boolean not null,
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_session;
-- +goose StatementEnd
