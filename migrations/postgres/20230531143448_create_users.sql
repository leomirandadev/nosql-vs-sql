-- +goose Up
-- +goose StatementBegin
create table users(
    id serial PRIMARY KEY,
    name varchar(80),
    city_id int,
    FOREIGN KEY (city_id) REFERENCES cities (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users
-- +goose StatementEnd
