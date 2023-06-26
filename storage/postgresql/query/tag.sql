-- name: UpsertTag :exec
insert into tag (id) values ($1)
on conflict (id) do nothing;

-- name: GetTag :one
SELECT * FROM tag
WHERE id = $1;