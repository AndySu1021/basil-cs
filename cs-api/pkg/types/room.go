package types

import (
	"database/sql"
	"time"
)

type RoomStatus int8

const (
	RoomStatusAll RoomStatus = iota
	RoomStatusPending
	RoomStatusQueuing
	RoomStatusServing
	RoomStatusClosed
)

type RoomSource int8

const (
	RoomSourceWeb RoomSource = iota + 1
	RoomSourceApp
)

type ListRoomParams struct {
	RoomID  int64      `form:"room_id" binding:""`
	StaffID int64      `form:"staff_id" binding:""`
	Status  RoomStatus `form:"status" binding:"min=0,max=3"`
	Pagination
}

type ListRoomRow struct {
	ID         int64        `db:"id" json:"id"`
	Status     RoomStatus   `db:"status" json:"status"`
	CreatedAt  time.Time    `db:"created_at" json:"created_at"`
	ClosedAt   sql.NullTime `db:"closed_at" json:"closed_at"`
	StaffName  string       `db:"staff_name" json:"staff_name"`
	MemberName string       `db:"member_name" json:"member_name"`
	TagName    string       `db:"tag_name" json:"tag_name"`
}

type FilterStaffRoomParams struct {
	StaffID *int64
}

type StaffRoom struct {
	ID           int64              `json:"id"`
	Status       RoomStatus         `json:"status"`
	MemberID     int64              `json:"member_id"`
	MemberName   string             `json:"member_name"`
	OnlineStatus MemberOnlineStatus `json:"online_status"`
	OpType       MessageOpType      `json:"op_type"`
	ContentType  MessageContentType `json:"content_type"`
	Content      string             `json:"content"`
}
