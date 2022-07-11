-- name: CreateTag :exec
INSERT INTO tag (name, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetTag :one
SELECT name, status
FROM tag
WHERE id = ? LIMIT 1;

-- name: UpdateTag :exec
UPDATE tag
SET name       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteTag :exec
DELETE
FROM tag
WHERE id = ?;

-- name: GetAllTag :many
select * from tag;

-- name: ListAvailableTag :many
select id, name
from tag
where status = 1
