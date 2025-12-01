-- +goose Up
-- +goose StatementBegin
create table if not exists habits (
    id uuid not null default gen_random_uuid() primary key,
    name varchar(300),
    user_id bigint not null,
    chat_id bigint not null,
    canceled boolean not null default false,
    created_at timestamp not null default current_timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists habits;
-- +goose StatementEnd
