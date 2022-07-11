// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package model

import (
	"context"
	"database/sql"
)

type Querier interface {
	AcceptRoom(ctx context.Context, arg AcceptRoomParams) error
	CheckFastReplyCategory(ctx context.Context, id int64) (interface{}, error)
	CloseRoom(ctx context.Context, arg CloseRoomParams) error
	ConstantSeeder(ctx context.Context) error
	CreateFAQ(ctx context.Context, arg CreateFAQParams) error
	CreateFastReply(ctx context.Context, arg CreateFastReplyParams) error
	CreateFastReplyCategory(ctx context.Context, arg CreateFastReplyCategoryParams) error
	CreateMember(ctx context.Context, arg CreateMemberParams) (sql.Result, error)
	CreateMessage(ctx context.Context, arg CreateMessageParams) error
	CreateNotice(ctx context.Context, arg CreateNoticeParams) error
	CreateRemind(ctx context.Context, arg CreateRemindParams) error
	CreateRole(ctx context.Context, arg CreateRoleParams) error
	CreateRoom(ctx context.Context, arg CreateRoomParams) (sql.Result, error)
	CreateSite(ctx context.Context, arg CreateSiteParams) error
	CreateStaff(ctx context.Context, arg CreateStaffParams) error
	CreateTag(ctx context.Context, arg CreateTagParams) error
	DeleteFAQ(ctx context.Context, id int64) error
	DeleteFastReply(ctx context.Context, id int64) error
	DeleteNotice(ctx context.Context, id int64) error
	DeleteRemind(ctx context.Context, id int64) error
	DeleteRole(ctx context.Context, id int64) error
	DeleteSite(ctx context.Context, id int64) error
	DeleteStaff(ctx context.Context, id int64) error
	DeleteTag(ctx context.Context, id int64) error
	GetAllAvailableFastReply(ctx context.Context) ([]GetAllAvailableFastReplyRow, error)
	GetAllRoles(ctx context.Context) ([]GetAllRolesRow, error)
	GetAllStaffs(ctx context.Context) ([]GetAllStaffsRow, error)
	GetAllTag(ctx context.Context) ([]Tag, error)
	GetCsConfig(ctx context.Context) (Constant, error)
	GetFAQ(ctx context.Context, id int64) (GetFAQRow, error)
	GetFastReply(ctx context.Context, id int64) (FastReply, error)
	GetGuestMember(ctx context.Context, arg GetGuestMemberParams) (Member, error)
	GetLastRoomMessage(ctx context.Context, roomID int64) (GetLastRoomMessageRow, error)
	GetLatestNotice(ctx context.Context) (GetLatestNoticeRow, error)
	// 獲取會員並未關閉的房間
	GetMemberAvailableRoom(ctx context.Context, memberID int64) (Room, error)
	GetNormalMember(ctx context.Context, arg GetNormalMemberParams) (Member, error)
	GetNotice(ctx context.Context, id int64) (Notice, error)
	GetOnlineStaffCount(ctx context.Context) (int64, error)
	GetQueuingRoomCount(ctx context.Context) (int64, error)
	GetRemind(ctx context.Context, id int64) (Remind, error)
	GetRole(ctx context.Context, id int64) (Role, error)
	GetRoom(ctx context.Context, id int64) (GetRoomRow, error)
	GetRoomInfo(ctx context.Context, id int64) (GetRoomInfoRow, error)
	GetServingRoomCount(ctx context.Context) (int64, error)
	GetSite(ctx context.Context, id int64) (GetSiteRow, error)
	GetSiteByCode(ctx context.Context, code string) (GetSiteByCodeRow, error)
	GetStaff(ctx context.Context, id int64) (GetStaffRow, error)
	GetStaffCountByRoleId(ctx context.Context, roleID int64) (int64, error)
	GetStaffServingRoomCount(ctx context.Context, staffID int64) (int64, error)
	GetTag(ctx context.Context, id int64) (GetTagRow, error)
	ListActiveRemind(ctx context.Context) ([]ListActiveRemindRow, error)
	ListAvailableFAQ(ctx context.Context) ([]ListAvailableFAQRow, error)
	ListAvailableStaff(ctx context.Context, id int64) ([]ListAvailableStaffRow, error)
	ListAvailableTag(ctx context.Context) ([]ListAvailableTagRow, error)
	ListFastReplyCategory(ctx context.Context) ([]Constant, error)
	ListMemberRoomMessage(ctx context.Context, arg ListMemberRoomMessageParams) ([]Message, error)
	ListQueuingRoom(ctx context.Context) ([]ListQueuingRoomRow, error)
	ListReportDailyStaff(ctx context.Context, arg ListReportDailyStaffParams) ([]ListReportDailyStaffRow, error)
	ListReportDailyTag(ctx context.Context, arg ListReportDailyTagParams) ([]ReportDailyTag, error)
	ListSite(ctx context.Context) ([]ListSiteRow, error)
	ListStaffRoomMessage(ctx context.Context, arg ListStaffRoomMessageParams) ([]Message, error)
	ListStaffServingRoom(ctx context.Context, staffID int64) ([]ListStaffServingRoomRow, error)
	RoleSeeder(ctx context.Context) error
	StaffLogin(ctx context.Context, arg StaffLoginParams) (StaffLoginRow, error)
	StaffSeeder(ctx context.Context, password string) error
	TagSeeder(ctx context.Context) error
	UpdateCsConfig(ctx context.Context, arg UpdateCsConfigParams) error
	UpdateFAQ(ctx context.Context, arg UpdateFAQParams) error
	UpdateFastReply(ctx context.Context, arg UpdateFastReplyParams) error
	UpdateMemberInfo(ctx context.Context, arg UpdateMemberInfoParams) error
	UpdateNotice(ctx context.Context, arg UpdateNoticeParams) error
	UpdateOnlineStatus(ctx context.Context, arg UpdateOnlineStatusParams) error
	UpdateRemind(ctx context.Context, arg UpdateRemindParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateRoomScore(ctx context.Context, arg UpdateRoomScoreParams) error
	UpdateRoomStaff(ctx context.Context, arg UpdateRoomStaffParams) error
	UpdateRoomStatus(ctx context.Context, arg UpdateRoomStatusParams) error
	UpdateSite(ctx context.Context, arg UpdateSiteParams) error
	UpdateStaff(ctx context.Context, arg UpdateStaffParams) error
	UpdateStaffAvatar(ctx context.Context, arg UpdateStaffAvatarParams) error
	UpdateStaffLogin(ctx context.Context, arg UpdateStaffLoginParams) error
	UpdateStaffLogout(ctx context.Context, arg UpdateStaffLogoutParams) error
	UpdateStaffServingStatus(ctx context.Context, arg UpdateStaffServingStatusParams) error
	UpdateStaffWithPassword(ctx context.Context, arg UpdateStaffWithPasswordParams) error
	UpdateTag(ctx context.Context, arg UpdateTagParams) error
}

var _ Querier = (*Queries)(nil)
