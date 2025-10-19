-- +goose Up
-- +goose StatementBegin
create table if not exists tasks (
    id uuid not null default gen_random_uuid() primary key,
    user_id bigint not null,
    is_done boolean not null default false,
    created_at timestamp not null default current_timestamp,

    name varchar(300),
    planed_to timestamp,
    planed_from timestamp,
    description varchar(500)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
