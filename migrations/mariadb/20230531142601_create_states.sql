-- +goose Up
-- +goose StatementBegin
create table states(
    id int auto_increment not null,
    name varchar(2),
    country varchar(40),
    PRIMARY KEY(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists states
-- +goose StatementEnd
