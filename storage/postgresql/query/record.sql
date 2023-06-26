-- name: CreateRecord :one
insert into record (id, content) values ($1, $2)
returning id;

-- name: GetRecordsByTag :many
select * from record join record_tag on id=record_id
where tag_id=$1;

-- name: AssociateNoteToTag :exec
insert into record_tag(record_id, tag_id) values ($1, $2);