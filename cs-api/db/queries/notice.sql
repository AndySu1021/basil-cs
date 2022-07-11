-- name: CreateNotice :exec
INSERT INTO notice (title, content, start_at, end_at, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetNotice :one
SELECT *
FROM notice
WHERE id = ? LIMIT 1;

-- name: UpdateNotice :exec
UPDATE notice
SET title      = ?,
    content    = ?,
    start_at   = ?,
    end_at     = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteNotice :exec
DELETE
FROM notice
WHERE id = ?;

-- name: GetLatestNotice :one
SELECT title, content
FROM notice
WHERE now() >= start_at
  AND now() <= end_at
  and status = 1
ORDER BY end_at
    LIMIT 1;
