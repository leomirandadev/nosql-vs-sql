-- +goose Up
-- +goose StatementBegin
create table users(
    id int auto_increment not null,
    name varchar(80),
    city_id int,
    PRIMARY KEY(id),
    FOREIGN KEY (city_id) REFERENCES cities (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users
-- +goose StatementEnd
