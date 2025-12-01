-- +goose Up
-- +goose StatementBegin
create table if not exists tasks (
    id uuid not null default gen_random_uuid() primary key,
    name varchar(300),
    user_id bigint not null,
    is_done boolean not null default false,
    planed_to timestamp,
    created_at timestamp not null default current_timestamp,
    planed_from timestamp,
    description varchar(500)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
