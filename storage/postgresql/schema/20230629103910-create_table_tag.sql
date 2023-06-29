
-- +migrate Up
create table if not exists tag (
  id varchar(128) primary key
);

-- +migrate Down
drop table if exists tag;