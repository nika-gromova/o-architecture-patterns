-- +goose Up
-- +goose StatementBegin
create table if not exists rules (

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
