-- +goose Up
-- +goose StatementBegin
create table states(
    id serial PRIMARY KEY,
    name varchar(2),
    country varchar(40)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists states
-- +goose StatementEnd
