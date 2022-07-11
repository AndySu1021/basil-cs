package notice

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc   iface.IAuthService
	noticeSvc iface.INoticeService
}

type CreateNoticeParams struct {
	Title   string       `json:"title" binding:"required"`
	Content string       `json:"content" binding:"required"`
	StartAt string       `json:"start_at" binding:"required"`
	EndAt   string       `json:"end_at" binding:"required"`
	Status  types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateNotice(c *gin.Context) {
	var (
		err           error
		requestParams CreateNoticeParams
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

	start, err := parseTime(requestParams.StartAt)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	end, err := parseTime(requestParams.EndAt)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params := model.CreateNoticeParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		StartAt:   start,
		EndAt:     end,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.noticeSvc.CreateNotice(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateNoticeParams struct {
	Title   string       `json:"title" binding:"required"`
	Content string       `json:"content" binding:"required"`
	StartAt string       `json:"start_at" binding:"required"`
	EndAt   string       `json:"end_at" binding:"required"`
	Status  types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateNotice(c *gin.Context) {
	var (
		err           error
		requestParams UpdateNoticeParams
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

	start, err := parseTime(requestParams.StartAt)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	end, err := parseTime(requestParams.EndAt)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params := model.UpdateNoticeParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		StartAt:   start,
		EndAt:     end,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.noticeSvc.UpdateNotice(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.noticeSvc.DeleteNotice(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	notice, err := h.noticeSvc.GetNotice(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, notice)
}

func (h *handler) ListNotice(c *gin.Context) {
	var (
		err    error
		params types.ListNoticeParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	notices, total, err := h.noticeSvc.ListNotice(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, notices, params.Pagination)
}

func (h *handler) GetLatestNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	notice, err := h.noticeSvc.GetLatestNotice(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ginTool.Success(c)
			return
		}
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, notice)
}

func parseTime(timeStr string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime.Add(-8 * time.Hour), nil
}

func NewHandler(authSvc iface.IAuthService, noticeSvc iface.INoticeService) *handler {
	return &handler{
		authSvc:   authSvc,
		noticeSvc: noticeSvc,
	}
}
