// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: room.sql

package model

import (
	"context"
	"database/sql"
	"time"

	"cs-api/pkg/types"
)

const acceptRoom = `-- name: AcceptRoom :exec
UPDATE room SET staff_id = ?, status = 3, accepted_at = ? WHERE id = ?
`

type AcceptRoomParams struct {
	StaffID    int64        `db:"staff_id" json:"staff_id"`
	AcceptedAt sql.NullTime `db:"accepted_at" json:"accepted_at"`
	ID         int64        `db:"id" json:"id"`
}

func (q *Queries) AcceptRoom(ctx context.Context, arg AcceptRoomParams) error {
	_, err := q.exec(ctx, q.acceptRoomStmt, acceptRoom, arg.StaffID, arg.AcceptedAt, arg.ID)
	return err
}

const closeRoom = `-- name: CloseRoom :exec
UPDATE room SET tag_id = ?, closed_at = ?, status = 4 WHERE id = ?
`

type CloseRoomParams struct {
	TagID    int64        `db:"tag_id" json:"tag_id"`
	ClosedAt sql.NullTime `db:"closed_at" json:"closed_at"`
	ID       int64        `db:"id" json:"id"`
}

func (q *Queries) CloseRoom(ctx context.Context, arg CloseRoomParams) error {
	_, err := q.exec(ctx, q.closeRoomStmt, closeRoom, arg.TagID, arg.ClosedAt, arg.ID)
	return err
}

const createRoom = `-- name: CreateRoom :execresult
INSERT INTO room (staff_id, member_id, ip, source, user_agent, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateRoomParams struct {
	StaffID   int64            `db:"staff_id" json:"staff_id"`
	MemberID  int64            `db:"member_id" json:"member_id"`
	Ip        string           `db:"ip" json:"ip"`
	Source    types.RoomSource `db:"source" json:"source"`
	UserAgent string           `db:"user_agent" json:"user_agent"`
	CreatedAt time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt time.Time        `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (sql.Result, error) {
	return q.exec(ctx, q.createRoomStmt, createRoom,
		arg.StaffID,
		arg.MemberID,
		arg.Ip,
		arg.Source,
		arg.UserAgent,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const getMemberAvailableRoom = `-- name: GetMemberAvailableRoom :one
SELECT id, staff_id, member_id, ip, source, user_agent, tag_id, score, status, accepted_at, closed_at, created_at, updated_at FROM room where member_id = ? and status <> 4 LIMIT 1
`

// 獲取會員並未關閉的房間
func (q *Queries) GetMemberAvailableRoom(ctx context.Context, memberID int64) (Room, error) {
	row := q.queryRow(ctx, q.getMemberAvailableRoomStmt, getMemberAvailableRoom, memberID)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.StaffID,
		&i.MemberID,
		&i.Ip,
		&i.Source,
		&i.UserAgent,
		&i.TagID,
		&i.Score,
		&i.Status,
		&i.AcceptedAt,
		&i.ClosedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getQueuingRoomCount = `-- name: GetQueuingRoomCount :one
SELECT count(*)
FROM room
WHERE status = 2
`

func (q *Queries) GetQueuingRoomCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getQueuingRoomCountStmt, getQueuingRoomCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getRoom = `-- name: GetRoom :one
SELECT room.id, room.staff_id, room.member_id, room.ip, room.source, room.user_agent, room.tag_id, room.score, room.status, room.accepted_at, room.closed_at, room.created_at, room.updated_at, member.site_id AS site_id, member.name AS member_name
FROM room
INNER JOIN member ON member.id = room.member_id
WHERE room.id = ? LIMIT 1
`

type GetRoomRow struct {
	ID         int64            `db:"id" json:"id"`
	StaffID    int64            `db:"staff_id" json:"staff_id"`
	MemberID   int64            `db:"member_id" json:"member_id"`
	Ip         string           `db:"ip" json:"ip"`
	Source     types.RoomSource `db:"source" json:"source"`
	UserAgent  string           `db:"user_agent" json:"user_agent"`
	TagID      int64            `db:"tag_id" json:"tag_id"`
	Score      int32            `db:"score" json:"score"`
	Status     types.RoomStatus `db:"status" json:"status"`
	AcceptedAt sql.NullTime     `db:"accepted_at" json:"accepted_at"`
	ClosedAt   sql.NullTime     `db:"closed_at" json:"closed_at"`
	CreatedAt  time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time        `db:"updated_at" json:"updated_at"`
	SiteID     int64            `db:"site_id" json:"site_id"`
	MemberName string           `db:"member_name" json:"member_name"`
}

func (q *Queries) GetRoom(ctx context.Context, id int64) (GetRoomRow, error) {
	row := q.queryRow(ctx, q.getRoomStmt, getRoom, id)
	var i GetRoomRow
	err := row.Scan(
		&i.ID,
		&i.StaffID,
		&i.MemberID,
		&i.Ip,
		&i.Source,
		&i.UserAgent,
		&i.TagID,
		&i.Score,
		&i.Status,
		&i.AcceptedAt,
		&i.ClosedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SiteID,
		&i.MemberName,
	)
	return i, err
}

const getRoomInfo = `-- name: GetRoomInfo :one
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
where r.id = ?
`

type GetRoomInfoRow struct {
	Ip           string                   `db:"ip" json:"ip"`
	Source       types.RoomSource         `db:"source" json:"source"`
	UserAgent    string                   `db:"user_agent" json:"user_agent"`
	MemberID     int64                    `db:"member_id" json:"member_id"`
	SiteName     string                   `db:"site_name" json:"site_name"`
	CreatedAt    time.Time                `db:"created_at" json:"created_at"`
	RealName     string                   `db:"real_name" json:"real_name"`
	Email        string                   `db:"email" json:"email"`
	Mobile       string                   `db:"mobile" json:"mobile"`
	Note         string                   `db:"note" json:"note"`
	OnlineStatus types.MemberOnlineStatus `db:"online_status" json:"online_status"`
}

func (q *Queries) GetRoomInfo(ctx context.Context, id int64) (GetRoomInfoRow, error) {
	row := q.queryRow(ctx, q.getRoomInfoStmt, getRoomInfo, id)
	var i GetRoomInfoRow
	err := row.Scan(
		&i.Ip,
		&i.Source,
		&i.UserAgent,
		&i.MemberID,
		&i.SiteName,
		&i.CreatedAt,
		&i.RealName,
		&i.Email,
		&i.Mobile,
		&i.Note,
		&i.OnlineStatus,
	)
	return i, err
}

const getServingRoomCount = `-- name: GetServingRoomCount :one
SELECT count(*)
FROM room
WHERE status = 3
`

func (q *Queries) GetServingRoomCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getServingRoomCountStmt, getServingRoomCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getStaffServingRoomCount = `-- name: GetStaffServingRoomCount :one
SELECT count(*) FROM room where staff_id = ? and status = 3
`

func (q *Queries) GetStaffServingRoomCount(ctx context.Context, staffID int64) (int64, error) {
	row := q.queryRow(ctx, q.getStaffServingRoomCountStmt, getStaffServingRoomCount, staffID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listQueuingRoom = `-- name: ListQueuingRoom :many
select room.id, room.status, room.member_id, member.name as member_name, member.online_status as online_status
from room
         inner join member on member.id = room.member_id
where room.status = 2 and member.online_status = 1
`

type ListQueuingRoomRow struct {
	ID           int64                    `db:"id" json:"id"`
	Status       types.RoomStatus         `db:"status" json:"status"`
	MemberID     int64                    `db:"member_id" json:"member_id"`
	MemberName   string                   `db:"member_name" json:"member_name"`
	OnlineStatus types.MemberOnlineStatus `db:"online_status" json:"online_status"`
}

func (q *Queries) ListQueuingRoom(ctx context.Context) ([]ListQueuingRoomRow, error) {
	rows, err := q.query(ctx, q.listQueuingRoomStmt, listQueuingRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListQueuingRoomRow{}
	for rows.Next() {
		var i ListQueuingRoomRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.MemberID,
			&i.MemberName,
			&i.OnlineStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStaffServingRoom = `-- name: ListStaffServingRoom :many
select room.id, room.status, room.member_id, member.name as member_name, member.online_status as online_status
from room
         inner join member on member.id = room.member_id
where staff_id = ?
  and room.status = 3
`

type ListStaffServingRoomRow struct {
	ID           int64                    `db:"id" json:"id"`
	Status       types.RoomStatus         `db:"status" json:"status"`
	MemberID     int64                    `db:"member_id" json:"member_id"`
	MemberName   string                   `db:"member_name" json:"member_name"`
	OnlineStatus types.MemberOnlineStatus `db:"online_status" json:"online_status"`
}

func (q *Queries) ListStaffServingRoom(ctx context.Context, staffID int64) ([]ListStaffServingRoomRow, error) {
	rows, err := q.query(ctx, q.listStaffServingRoomStmt, listStaffServingRoom, staffID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListStaffServingRoomRow{}
	for rows.Next() {
		var i ListStaffServingRoomRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.MemberID,
			&i.MemberName,
			&i.OnlineStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoomScore = `-- name: UpdateRoomScore :exec
UPDATE room SET score = ? WHERE id = ? and status = 3
`

type UpdateRoomScoreParams struct {
	Score int32 `db:"score" json:"score"`
	ID    int64 `db:"id" json:"id"`
}

func (q *Queries) UpdateRoomScore(ctx context.Context, arg UpdateRoomScoreParams) error {
	_, err := q.exec(ctx, q.updateRoomScoreStmt, updateRoomScore, arg.Score, arg.ID)
	return err
}

const updateRoomStaff = `-- name: UpdateRoomStaff :exec
update room
set staff_id = ?
where id = ?
`

type UpdateRoomStaffParams struct {
	StaffID int64 `db:"staff_id" json:"staff_id"`
	ID      int64 `db:"id" json:"id"`
}

func (q *Queries) UpdateRoomStaff(ctx context.Context, arg UpdateRoomStaffParams) error {
	_, err := q.exec(ctx, q.updateRoomStaffStmt, updateRoomStaff, arg.StaffID, arg.ID)
	return err
}

const updateRoomStatus = `-- name: UpdateRoomStatus :exec
UPDATE room
SET status = ?
where id = ?
`

type UpdateRoomStatusParams struct {
	Status types.RoomStatus `db:"status" json:"status"`
	ID     int64            `db:"id" json:"id"`
}

func (q *Queries) UpdateRoomStatus(ctx context.Context, arg UpdateRoomStatusParams) error {
	_, err := q.exec(ctx, q.updateRoomStatusStmt, updateRoomStatus, arg.Status, arg.ID)
	return err
}
