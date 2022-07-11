-- name: CreateRoom :execresult
INSERT INTO room (staff_id, member_id, ip, source, user_agent, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- 獲取會員並未關閉的房間
-- name: GetMemberAvailableRoom :one
SELECT * FROM room where member_id = ? and status <> 4 LIMIT 1;

-- name: GetRoom :one
SELECT room.*, member.site_id AS site_id, member.name AS member_name
FROM room
INNER JOIN member ON member.id = room.member_id
WHERE room.id = ? LIMIT 1;

-- name: AcceptRoom :exec
UPDATE room SET staff_id = ?, status = 3, accepted_at = ? WHERE id = ?;

-- name: CloseRoom :exec
UPDATE room SET tag_id = ?, closed_at = ?, status = 4 WHERE id = ?;

-- name: UpdateRoomScore :exec
UPDATE room SET score = ? WHERE id = ? and status = 3;

-- name: ListStaffServingRoom :many
select room.id, room.status, room.member_id, member.name as member_name, member.online_status as online_status
from room
         inner join member on member.id = room.member_id
where staff_id = ?
  and room.status = 3;

-- name: ListQueuingRoom :many
select room.id, room.status, room.member_id, member.name as member_name, member.online_status as online_status
from room
         inner join member on member.id = room.member_id
where room.status = 2 and member.online_status = 1;

-- name: GetStaffServingRoomCount :one
SELECT count(*) FROM room where staff_id = ? and status = 3;

-- name: UpdateRoomStaff :exec
update room
set staff_id = ?
where id = ?;

-- name: GetRoomInfo :one
select r.ip,
       r.source,
       r.user_agent,
       r.member_id,
       s.name AS site_name,
       m.created_at,
       m.real_name,
       m.email,
       m.mobile,
       m.note,
       m.online_status
from room r
         inner join member m on m.id = r.member_id
         inner join site s on s.id = m.site_id
where r.id = ?;

-- name: GetQueuingRoomCount :one
SELECT count(*)
FROM room
WHERE status = 2;

-- name: GetServingRoomCount :one
SELECT count(*)
FROM room
WHERE status = 3;

-- name: UpdateRoomStatus :exec
UPDATE room
SET status = ?
where id = ?;