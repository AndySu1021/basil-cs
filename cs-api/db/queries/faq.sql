-- name: CreateFAQ :exec
INSERT INTO faq (question, answer, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetFAQ :one
SELECT question, answer, status
FROM faq
WHERE id = ? LIMIT 1;

-- name: UpdateFAQ :exec
UPDATE faq
SET question   = ?,
    answer     = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteFAQ :exec
DELETE
FROM faq
WHERE id = ?;

-- name: ListAvailableFAQ :many
select question, answer
from faq
where status = 1
