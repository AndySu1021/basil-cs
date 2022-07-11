-- name: UpdateCsConfig :exec
UPDATE constant
SET value      = ?,
    updated_by = ?,
    updated_at = ?
WHERE `type` = 2
  and `key` = 'CsConfig';

-- name: GetCsConfig :one
SELECT *
FROM constant
WHERE `type` = 2
  and `key` = 'CsConfig' LIMIT 1;

-- name: CreateFastReplyCategory :exec
INSERT INTO constant (`type`, `key`, value, created_by, created_at, updated_by, updated_at)
VALUES (1, 'FastReply', ?, ?, ?, ?, ?);

-- name: ListFastReplyCategory :many
SELECT *
FROM constant
WHERE `type` = 1
  and `key` = 'FastReply';

-- name: CheckFastReplyCategory :one
SELECT 1
FROM constant
WHERE id = ? LIMIT 1;
