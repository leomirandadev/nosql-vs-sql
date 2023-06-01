-- +goose Up
-- +goose StatementBegin
create table cities(
    id int auto_increment not null,
    name varchar(40),
    state_id int,
    PRIMARY KEY(id),
    FOREIGN KEY (state_id) REFERENCES states (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cities
-- +goose StatementEnd
