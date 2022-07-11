package message

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	authSvc    iface.IAuthService
	messageSvc iface.IMessageService
}

func (h *handler) ListMessage(c *gin.Context) {
	var (
		err    error
		params types.ListMessageParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	messages, total, err := h.messageSvc.ListMessage(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, messages, params.Pagination)
}

type ListStaffRoomMessageParams struct {
	RoomID int64 `form:"room_id" binding:"required,gte=1"`
	types.Pagination
}

func (h *handler) ListStaffRoomMessage(c *gin.Context) {
	var (
		err           error
		requestParams ListStaffRoomMessageParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := model.ListStaffRoomMessageParams{
		RoomID: requestParams.RoomID,
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	if requestParams.Pagination.Page > 0 {
		params.Offset = (requestParams.Pagination.Page - 1) * requestParams.Pagination.PageSize
	}

	messages, err := h.messageSvc.ListRoomMessage(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, messages)
}

type ListMemberRoomMessageParams struct {
	types.Pagination
}

func (h *handler) ListMemberRoomMessage(c *gin.Context) {
	var (
		err           error
		requestParams ListMemberRoomMessageParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	clientInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeMember)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params := model.ListMemberRoomMessageParams{
		RoomID: clientInfo.RoomID,
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	if requestParams.Pagination.Page > 0 {
		params.Offset = (requestParams.Pagination.Page - 1) * requestParams.Pagination.PageSize
	}

	messages, err := h.messageSvc.ListRoomMessage(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, messages)
}

func NewHandler(authSvc iface.IAuthService, messageSvc iface.IMessageService) *handler {
	return &handler{
		authSvc:    authSvc,
		messageSvc: messageSvc,
	}
}
