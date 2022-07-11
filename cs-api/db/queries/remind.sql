-- name: CreateRemind :exec
INSERT INTO remind (title, content, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetRemind :one
SELECT *
FROM remind
WHERE id = ? LIMIT 1;

-- name: UpdateRemind :exec
UPDATE remind
SET title      = ?,
    content    = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteRemind :exec
DELETE
FROM remind
WHERE id = ?;

-- name: ListActiveRemind :many
SELECT title, content
FROM remind
WHERE status = 1