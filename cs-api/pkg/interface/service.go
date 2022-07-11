package iface

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"github.com/gin-gonic/gin"
	"time"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (pkg.ClientInfo, error)
	GetStaffInfo(ctx context.Context, staffId int64) (model.GetStaffRow, error)
	Logout(ctx context.Context, userInfo pkg.ClientInfo) error
	SetClientInfo(clientType pkg.ClientType) gin.HandlerFunc
	GetClientInfo(ctx context.Context, clientType pkg.ClientType) (pkg.ClientInfo, error)
	CheckPermission(permission string) gin.HandlerFunc
}

type IStaffService interface {
	ListStaff(ctx context.Context, params types.ListStaffParams) ([]types.ListStaffRow, int64, error)
	GetStaff(ctx context.Context, staffId int64) (model.GetStaffRow, error)
	CreateStaff(ctx context.Context, params model.CreateStaffParams) error
	UpdateStaff(ctx context.Context, params interface{}) error
	DeleteStaff(ctx context.Context, staffId int64) error
	UpdateStaffServingStatus(ctx context.Context, staffId int64, status types.StaffServingStatus) error
	ListAvailableStaff(ctx context.Context, staffId int64) ([]model.ListAvailableStaffRow, error)
	GetAllStaffs(ctx context.Context) ([]model.GetAllStaffsRow, error)
}

type IRoomService interface {
	CreateRoom(ctx context.Context, source types.RoomSource, userAgent string, memberId int64) (int64, error)
	AcceptRoom(ctx context.Context, staffId int64, roomId int64) error
	CloseRoom(ctx context.Context, staffId, roomId, tagId int64) error
	TransferRoom(ctx context.Context, staffId, roomId, toStaffId int64) error
	UpdateRoomScore(ctx context.Context, roomId int64, score int32) error
	ListStaffServingRoom(ctx context.Context, staffId int64) ([]types.StaffRoom, error)
	ListQueuingRoom(ctx context.Context) ([]model.ListQueuingRoomRow, error)
	ListRoom(ctx context.Context, params types.ListRoomParams) ([]types.ListRoomRow, int64, error)
	GetStaffServingRoomCount(ctx context.Context, staffId int64) (int64, error)
	GetMemberAvailableRoom(ctx context.Context, memberId int64) (model.Room, error)
	GetRoomInfo(ctx context.Context, roomId int64) (model.GetRoomInfoRow, error)
	UpdateRoomStatus(ctx context.Context, params model.UpdateRoomStatusParams) error
}

type IMessageService interface {
	CreateMessage(ctx context.Context, params model.CreateMessageParams) error
	ListRoomMessage(ctx context.Context, params interface{}) ([]model.Message, error)
	ListMessage(ctx context.Context, params types.ListMessageParams) ([]types.ListMessageRow, int64, error)
}

type ITagService interface {
	ListTag(ctx context.Context, params types.ListTagParams) ([]types.ListTagRow, int64, error)
	GetTag(ctx context.Context, tagId int64) (model.GetTagRow, error)
	CreateTag(ctx context.Context, params model.CreateTagParams) error
	UpdateTag(ctx context.Context, params model.UpdateTagParams) error
	DeleteTag(ctx context.Context, tagId int64) error
	ListAvailableTag(ctx context.Context) ([]model.ListAvailableTagRow, error)
}

type IFastReplyService interface {
	ListFastReply(ctx context.Context, params types.ListFastReplyParams) ([]types.ListFastReplyRow, int64, error)
	GetFastReply(ctx context.Context, id int64) (model.FastReply, error)
	CreateFastReply(ctx context.Context, params model.CreateFastReplyParams) error
	UpdateFastReply(ctx context.Context, params model.UpdateFastReplyParams) error
	DeleteFastReply(ctx context.Context, id int64) error
	ListCategory(ctx context.Context) ([]model.Constant, error)
	CreateCategory(ctx context.Context, params model.CreateFastReplyCategoryParams) error
	CheckCategory(ctx context.Context, id int64) (interface{}, error)
	ListFastReplyGroup(ctx context.Context) ([]pkg.FastReplyGroupItem, error)
}

type IReportService interface {
	ListDailyTagReport(ctx context.Context, startDate, endDate time.Time) ([]pkg.DailyTagReportColumn, map[string]map[string]int32, error)
	ListDailyStaffReport(ctx context.Context, startDate, endDate time.Time) ([]model.ListReportDailyStaffRow, error)
}

type ICsConfigService interface {
	GetCsConfig(ctx context.Context) (types.CsConfig, error)
	UpdateCsConfig(ctx context.Context, staffId int64, config types.CsConfig) error
}

type IMemberService interface {
	ListMember(ctx context.Context, params types.ListMemberParams) ([]types.ListMemberRow, int64, error)
	GetOrCreateMember(ctx context.Context, name string, deviceId string, siteId int64) (model.Member, error)
	UpdateOnlineStatus(ctx context.Context, params model.UpdateOnlineStatusParams) error
	UpdateMemberInfo(ctx context.Context, params model.UpdateMemberInfoParams) error
}

type IRoleService interface {
	ListRole(ctx context.Context, params types.ListRoleParams) ([]types.ListRoleRow, int64, error)
	GetAllRoles(ctx context.Context) ([]model.GetAllRolesRow, error)
	GetRole(ctx context.Context, roleId int64) (model.Role, error)
	CreateRole(ctx context.Context, params model.CreateRoleParams) error
	UpdateRole(ctx context.Context, params model.UpdateRoleParams) error
	DeleteRole(ctx context.Context, roleId int64) error
}

// INoticeService 會員通知訊息
type INoticeService interface {
	ListNotice(ctx context.Context, params types.ListNoticeParams) ([]types.ListNoticeRow, int64, error)
	GetNotice(ctx context.Context, noticeId int64) (model.Notice, error)
	CreateNotice(ctx context.Context, params model.CreateNoticeParams) error
	UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) error
	DeleteNotice(ctx context.Context, noticeId int64) error
	GetLatestNotice(ctx context.Context) (model.GetLatestNoticeRow, error)
}

type IRemindService interface {
	ListRemind(ctx context.Context, params types.ListRemindParams) ([]types.ListRemindRow, int64, error)
	GetRemind(ctx context.Context, remindId int64) (model.Remind, error)
	CreateRemind(ctx context.Context, params model.CreateRemindParams) error
	UpdateRemind(ctx context.Context, params model.UpdateRemindParams) error
	DeleteRemind(ctx context.Context, remindId int64) error
	ListActiveRemind(ctx context.Context) ([]model.ListActiveRemindRow, error)
}

type IFAQService interface {
	ListFAQ(ctx context.Context, params types.ListFAQParams) ([]types.ListFAQRow, int64, error)
	GetFAQ(ctx context.Context, faqId int64) (model.GetFAQRow, error)
	CreateFAQ(ctx context.Context, params model.CreateFAQParams) error
	UpdateFAQ(ctx context.Context, params model.UpdateFAQParams) error
	DeleteFAQ(ctx context.Context, faqId int64) error
	ListAvailableFAQ(ctx context.Context) ([]model.ListAvailableFAQRow, error)
}

type ISiteService interface {
	ListSite(ctx context.Context) ([]model.ListSiteRow, error)
	GetSite(ctx context.Context, siteId int64) (model.GetSiteRow, error)
	CreateSite(ctx context.Context, params model.CreateSiteParams) error
	UpdateSite(ctx context.Context, params model.UpdateSiteParams) error
	DeleteSite(ctx context.Context, siteId int64) error
	GetSiteByCode(ctx context.Context, code string) (model.GetSiteByCodeRow, error)
}

type ICommonService interface {
	GetPanelData(ctx context.Context) (types.PanelData, error)
}
