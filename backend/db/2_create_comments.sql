-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS comments (
  id                     serial primary key,
  post_id                INTEGER REFERENCES posts(id),
  body                   TEXT,
  created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE comments;
