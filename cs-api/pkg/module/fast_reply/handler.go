package fast_reply

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc      iface.IAuthService
	fastReplySvc iface.IFastReplyService
}

type CreateFastReplyParams struct {
	CategoryID int64        `json:"category_id" binding:"required"`
	Title      string       `json:"title" binding:"required"`
	Content    string       `json:"content" binding:"required"`
	Status     types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateFastReply(c *gin.Context) {
	var (
		err           error
		requestParams CreateFastReplyParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	now := time.Now().UTC()
	params := model.CreateFastReplyParams{
		CategoryID: requestParams.CategoryID,
		Title:      requestParams.Title,
		Content:    requestParams.Content,
		Status:     requestParams.Status,
		CreatedBy:  staffInfo.ID,
		CreatedAt:  now,
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  now,
	}

	_, err = h.fastReplySvc.CheckCategory(ctx, requestParams.CategoryID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.fastReplySvc.CreateFastReply(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateFastReplyParams struct {
	CategoryID int64        `json:"category_id" binding:"required"`
	Title      string       `json:"title" binding:"required"`
	Content    string       `json:"content" binding:"required"`
	Status     types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateFastReply(c *gin.Context) {
	var (
		err           error
		requestParams UpdateFastReplyParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := model.UpdateFastReplyParams{
		CategoryID: requestParams.CategoryID,
		Title:      requestParams.Title,
		Content:    requestParams.Content,
		Status:     requestParams.Status,
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  time.Now().UTC(),
		ID:         id,
	}

	if err = h.fastReplySvc.UpdateFastReply(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteFastReply(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.fastReplySvc.DeleteFastReply(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetFastReply(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	fastReply, err := h.fastReplySvc.GetFastReply(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, fastReply)
}

func (h *handler) ListFastReply(c *gin.Context) {
	var (
		err    error
		params types.ListFastReplyParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	fastReplies, total, err := h.fastReplySvc.ListFastReply(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, fastReplies, params.Pagination)
}

func (h *handler) ListFastReplyGroup(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	group, err := h.fastReplySvc.ListFastReplyGroup(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, group)
}

func (h *handler) ListFastReplyCategory(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	categories, err := h.fastReplySvc.ListCategory(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, categories)
}

type CreateFastReplyCategoryParams struct {
	Name string `json:"name" binding:"required"`
}

func (h *handler) CreateFastReplyCategory(c *gin.Context) {
	var (
		err           error
		requestParams CreateFastReplyCategoryParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	now := time.Now().UTC()
	params := model.CreateFastReplyCategoryParams{
		Value:     requestParams.Name,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.fastReplySvc.CreateCategory(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func NewHandler(authSvc iface.IAuthService, fastReplySvc iface.IFastReplyService) *handler {
	return &handler{
		authSvc:      authSvc,
		fastReplySvc: fastReplySvc,
	}
}
