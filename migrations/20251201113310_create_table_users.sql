-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id uuid not null default gen_random_uuid() primary key,
    name varchar(50),
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user;
-- +goose StatementEnd
