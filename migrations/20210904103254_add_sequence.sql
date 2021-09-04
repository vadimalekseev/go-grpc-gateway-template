-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE links_id_seq
    START WITH 1
    INCREMENT BY 1;
ALTER TABLE links
    ALTER COLUMN id SET DEFAULT nextval('links_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE links
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE links_id_seq;
-- +goose StatementEnd
