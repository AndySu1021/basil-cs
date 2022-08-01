package room

import (
	"context"
	"crypto/md5"
	"cs-api/config"
	"cs-api/internal"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"database/sql"
	"encoding/json"
	"fmt"
	errorTool "github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"strconv"
	"time"
)

type handler struct {
	config    *config.Config
	authSvc   iface.IAuthService
	siteSvc   iface.ISiteService
	memberSvc iface.IMemberService
	roomSvc   iface.IRoomService
	lua       iface.ILuaScript
}

func (h *handler) ListRoom(c *gin.Context) {
	var (
		err    error
		params types.ListRoomParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	rooms, total, err := h.roomSvc.ListRoom(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, rooms, params.Pagination)
}

func (h *handler) ListStaffServingRoom(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	rooms, err := h.roomSvc.ListStaffServingRoom(ctx, staffInfo.ID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, rooms)
}

func (h *handler) ListQueuingRoom(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	rooms, err := h.roomSvc.ListQueuingRoom(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, rooms)
}

func (h *handler) AcceptRoom(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roomSvc.AcceptRoom(ctx, staffInfo.ID, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type CloseRoomParams struct {
	TagID int64 `json:"tag_id" binding:"required"`
}

func (h *handler) CloseRoom(c *gin.Context) {
	var (
		err           error
		requestParams CloseRoomParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roomSvc.CloseRoom(ctx, staffInfo.ID, id, requestParams.TagID); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type TransferRoomParams struct {
	StaffID int64 `json:"staff_id" binding:"required"`
}

func (h *handler) TransferRoom(c *gin.Context) {
	var (
		err           error
		requestParams TransferRoomParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roomSvc.TransferRoom(ctx, staffInfo.ID, id, requestParams.StaffID); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateRoomScoreParams struct {
	Score int32 `json:"score" binding:"required,min=1,max=5"`
}

func (h *handler) UpdateRoomScore(c *gin.Context) {
	var (
		err           error
		requestParams UpdateRoomScoreParams
		ctx           = c.Request.Context()
	)

	memberInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeMember)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	if err = h.roomSvc.UpdateRoomScore(ctx, memberInfo.RoomID, requestParams.Score); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type CreateRoomParams struct {
	Code   string           `json:"code" binding:"required"`
	Source types.RoomSource `json:"source" binding:"required,oneof=1 2"`
	Name   string           `json:"name" binding:""`
}

func (h *handler) CreateRoom(c *gin.Context) {
	var (
		err       error
		params    CreateRoomParams
		ctx       = c.Request.Context()
		userAgent = c.GetHeader("User-Agent")
	)

	if err = c.ShouldBindJSON(&params); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	site, err := h.siteSvc.GetSiteByCode(ctx, params.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("no matching code")
		}
		ginTool.Error(c, err)
		return
	}

	deviceId := getDeviceId(ctx, params.Code, params.Source, userAgent)
	member, err := h.memberSvc.GetOrCreateMember(ctx, params.Name, deviceId, site.ID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	roomId, err := h.roomSvc.CreateRoom(ctx, params.Source, userAgent, member.ID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	name := fmt.Sprintf("%d-%s", site.ID, member.Name)

	if err = h.lua.RemoveMemberToken(ctx, name); err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	token := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+name)))

	clientInfo := internal.ClientInfo{
		ID:     member.ID,
		Type:   internal.ClientTypeMember,
		Name:   member.Name,
		RoomID: roomId,
		SiteID: site.ID,
	}
	payload, _ := json.Marshal(clientInfo)
	if err = h.lua.SetToken(ctx, "member", name, token, payload, 1*time.Hour); err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	ginTool.SuccessWithData(c, fmt.Sprintf("%s?room_id=%d&name=%s&sid=%s", h.config.WebChatBaseUrl, roomId, member.Name, token))
}

func getDeviceId(ctx context.Context, code string, source types.RoomSource, userAgent string) string {
	ip := ctx.Value(ginTool.ContextKeyClientIp).(string)
	return fmt.Sprintf("%x", md5.Sum([]byte(code+strconv.Itoa(int(source))+userAgent+ip)))
}

func (h *handler) GetRoomInfo(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	info, err := h.roomSvc.GetRoomInfo(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, info)
}

type HandlerParams struct {
	fx.In

	Config    *config.Config
	AuthSvc   iface.IAuthService
	SiteSvc   iface.ISiteService
	MemberSvc iface.IMemberService
	RoomSvc   iface.IRoomService
	Lua       iface.ILuaScript
}

func NewHandler(p HandlerParams) *handler {
	return &handler{
		config:    p.Config,
		authSvc:   p.AuthSvc,
		siteSvc:   p.SiteSvc,
		memberSvc: p.MemberSvc,
		roomSvc:   p.RoomSvc,
		lua:       p.Lua,
	}
}
