-- name: CreateSite :exec
INSERT INTO site (name, code, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetSite :one
SELECT name, code, status
FROM site
WHERE id = ? LIMIT 1;

-- name: UpdateSite :exec
UPDATE site
SET name       = ?,
    code     = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteSite :exec
DELETE
FROM site
WHERE id = ?;

-- name: ListSite :many
select id, name, code, status
from site;

-- name: GetSiteByCode :one
select id, name, status
from site
where code = ?
  and status = 1;
