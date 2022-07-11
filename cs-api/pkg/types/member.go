package types

import (
	"time"
)

type MemberType int8

const (
	MemberTypeNormal MemberType = iota + 1
	MemberTypeGuest
)

type MemberOnlineStatus int8

const (
	MemberOnlineStatusOnline MemberOnlineStatus = iota + 1
	MemberOnlineStatusOffline
)

type ListMemberParams struct {
	Mobile string `form:"mobile" binding:""`
	Email  string `form:"email" binding:""`
	Pagination
}

type ListMemberRow struct {
	ID           int64              `db:"id" json:"id"`
	Type         MemberType         `db:"type" json:"type"`
	Name         string             `db:"name" json:"name"`
	RealName     string             `db:"real_name" json:"real_name"`
	Mobile       string             `db:"mobile" json:"mobile"`
	Email        string             `db:"email" json:"email"`
	Status       Status             `db:"status" json:"status"`
	OnlineStatus MemberOnlineStatus `db:"online_status" json:"online_status"`
	CreatedAt    time.Time          `db:"created_at" json:"created_at"`
	SiteName     string             `db:"site_name" json:"site_name"`
}
