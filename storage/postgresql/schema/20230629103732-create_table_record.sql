
-- +migrate Up
create table if not exists record (
  id uuid primary key,
  content text not null
);

-- +migrate Down
drop table if exists record;