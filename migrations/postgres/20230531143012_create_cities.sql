-- +goose Up
-- +goose StatementBegin
create table cities(
    id serial PRIMARY KEY,
    name varchar(40),
    state_id int,
    FOREIGN KEY (state_id) REFERENCES states (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cities
-- +goose StatementEnd
