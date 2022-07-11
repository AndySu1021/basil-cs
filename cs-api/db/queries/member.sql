-- name: GetGuestMember :one
select *
from member
where type = 2
  and device_id = ?
  and site_id = ? LIMIT 1;

-- name: GetNormalMember :one
select *
from member
where type = 1
  and site_id = ?
  and name = ? LIMIT 1;

-- name: CreateMember :execresult
INSERT INTO member (type, site_id, name, device_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateOnlineStatus :exec
update member
set online_status = ?
where id = ?;

-- name: UpdateMemberInfo :exec
update member
set real_name  = ?,
    email      = ?,
    mobile     = ?,
    note       = ?,
    updated_at = ?
where id = ?;