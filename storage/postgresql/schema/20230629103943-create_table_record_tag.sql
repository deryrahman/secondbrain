
-- +migrate Up
create table if not exists record_tag (
  record_id uuid,
  tag_id varchar(128),
  primary key (record_id, tag_id),
  constraint fk_record foreign key (record_id) references record(id),
  constraint fk_tag foreign key (tag_id) references tag(id)
);

-- +migrate Down
drop table if exists record_tag;