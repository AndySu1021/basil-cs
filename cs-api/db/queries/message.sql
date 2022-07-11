-- name: CreateMessage :exec
INSERT INTO message (room_id, op_type, sender_type, sender_id, sender_name, content_type, content, extra, ts, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListStaffRoomMessage :many
select *
from message
where room_id = ?
order by ts desc
limit ? offset ?;

-- name: ListMemberRoomMessage :many
select *
from message
where room_id = ?
  and sender_type <> 0
order by ts desc
limit ? offset ?;

-- name: GetLastRoomMessage :one
select op_type, content_type, content
from message
where room_id = ?
order by ts desc
LIMIT 1;