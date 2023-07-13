-- name: CreateRecord :one
insert into record (id, content) values ($1, $2)
returning id;

-- name: GetRecordsByTag :many
with records_tags as (
  select distinct(id), content from record join record_tag on id=record_id where tag_id=any($1::text[])
)
select r.id, r.content, array_agg(rt.tag_id)::text[] as tags
from records_tags as r join record_tag as rt on r.id=rt.record_id
group by r.id, r.content;

-- name: AssociateNoteToTag :exec
insert into record_tag(record_id, tag_id) values ($1, $2);