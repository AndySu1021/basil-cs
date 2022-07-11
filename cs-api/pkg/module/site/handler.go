package site

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
	authSvc iface.IAuthService
	siteSvc iface.ISiteService
}

type CreateSiteParams struct {
	Name   string       `json:"name" binding:"required"`
	Code   string       `json:"code" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateSite(c *gin.Context) {
	var (
		err           error
		requestParams CreateSiteParams
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
	params := model.CreateSiteParams{
		Name:      requestParams.Name,
		Code:      requestParams.Code,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.siteSvc.CreateSite(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateSiteParams struct {
	Name   string       `json:"name" binding:"required"`
	Code   string       `json:"code" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateSite(c *gin.Context) {
	var (
		err           error
		requestParams UpdateSiteParams
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

	params := model.UpdateSiteParams{
		Name:      requestParams.Name,
		Code:      requestParams.Code,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.siteSvc.UpdateSite(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteSite(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.siteSvc.DeleteSite(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) ListSite(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	sites, err := h.siteSvc.ListSite(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, sites)
}

func (h *handler) GetSite(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	merchant, err := h.siteSvc.GetSite(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, merchant)
}

func NewHandler(authSvc iface.IAuthService, siteSvc iface.ISiteService) *handler {
	return &handler{
		authSvc: authSvc,
		siteSvc: siteSvc,
	}
}
