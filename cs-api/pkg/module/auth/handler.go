package auth

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"database/sql"
	"fmt"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	authSvc iface.IAuthService
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *handler) Login(c *gin.Context) {
	var err error

	var params LoginParams
	if err = c.ShouldBindJSON(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	ctx := c.Request.Context()
	staffInfo, err := h.authSvc.Login(ctx, params.Username, params.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			ginTool.Error(c, fmt.Errorf("帳號或密碼錯誤"))
			return
		}
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{
		"staffID":     staffInfo.ID,
		"username":    staffInfo.Username,
		"token":       staffInfo.Token,
		"permissions": staffInfo.Permissions,
	})
}

func (h *handler) GetStaffInfo(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	info, err := h.authSvc.GetStaffInfo(ctx, staffInfo.ID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, info)
}

func (h *handler) Logout(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.authSvc.Logout(ctx, staffInfo); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func NewHandler(authSvc iface.IAuthService) *handler {
	return &handler{
		authSvc: authSvc,
	}
}
