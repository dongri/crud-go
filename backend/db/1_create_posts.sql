-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS posts (
  id                     serial primary key,
  title                  VARCHAR(256),
  body                   TEXT,
  created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE posts;
