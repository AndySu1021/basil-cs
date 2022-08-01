package role

import (
	"cs-api/db/model"
	"cs-api/internal"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"encoding/json"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc iface.IAuthService
	roleSvc iface.IRoleService
}

type CreateRoleParams struct {
	Name        string   `json:"name" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}

func (h *handler) CreateRole(c *gin.Context) {
	var (
		err           error
		requestParams CreateRoleParams
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
	result, _ := json.Marshal(requestParams.Permissions)
	params := model.CreateRoleParams{
		Name:        requestParams.Name,
		Permissions: result,
		CreatedBy:   staffInfo.ID,
		CreatedAt:   now,
		UpdatedBy:   staffInfo.ID,
		UpdatedAt:   now,
	}

	if err = h.roleSvc.CreateRole(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateRoleParams struct {
	Name        string   `json:"name" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}

func (h *handler) UpdateRole(c *gin.Context) {
	var (
		err           error
		requestParams UpdateRoleParams
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

	result, _ := json.Marshal(requestParams.Permissions)
	params := model.UpdateRoleParams{
		Name:        requestParams.Name,
		Permissions: result,
		UpdatedBy:   staffInfo.ID,
		UpdatedAt:   time.Now().UTC(),
		ID:          id,
	}

	if err = h.roleSvc.UpdateRole(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteRole(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roleSvc.DeleteRole(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetRole(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	role, err := h.roleSvc.GetRole(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, role)
}

func (h *handler) ListRole(c *gin.Context) {
	var (
		err    error
		params types.ListRoleParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	roles, total, err := h.roleSvc.ListRole(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, roles, params.Pagination)
}

func (h *handler) GetAllRoles(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	roles, err := h.roleSvc.GetAllRoles(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, roles)
}

func NewHandler(authSvc iface.IAuthService, roleSvc iface.IRoleService) *handler {
	return &handler{
		authSvc: authSvc,
		roleSvc: roleSvc,
	}
}
