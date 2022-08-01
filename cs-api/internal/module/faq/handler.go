package faq

import (
	"cs-api/db/model"
	"cs-api/internal"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc iface.IAuthService
	faqSvc  iface.IFAQService
}

type CreateFAQParams struct {
	Question string       `json:"question" binding:"required"`
	Answer   string       `json:"answer" binding:"required"`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateFAQ(c *gin.Context) {
	var (
		err           error
		requestParams CreateFAQParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, internal.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	now := time.Now().UTC()
	params := model.CreateFAQParams{
		Question:  requestParams.Question,
		Answer:    requestParams.Answer,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.faqSvc.CreateFAQ(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateFAQParams struct {
	Question string       `json:"question" binding:"required"`
	Answer   string       `json:"answer" binding:"required"`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateFAQ(c *gin.Context) {
	var (
		err           error
		requestParams UpdateFAQParams
		ctx           = c.Request.Context()
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

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := model.UpdateFAQParams{
		Question:  requestParams.Question,
		Answer:    requestParams.Answer,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.faqSvc.UpdateFAQ(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteFAQ(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.faqSvc.DeleteFAQ(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) ListAvailableFAQ(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	tags, err := h.faqSvc.ListAvailableFAQ(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tags)
}

func (h *handler) GetFAQ(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tag, err := h.faqSvc.GetFAQ(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tag)
}

func (h *handler) ListFAQ(c *gin.Context) {
	var (
		err    error
		params types.ListFAQParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	faqs, total, err := h.faqSvc.ListFAQ(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, faqs, params.Pagination)
}

func NewHandler(authSvc iface.IAuthService, faqSvc iface.IFAQService) *handler {
	return &handler{
		authSvc: authSvc,
		faqSvc:  faqSvc,
	}
}
