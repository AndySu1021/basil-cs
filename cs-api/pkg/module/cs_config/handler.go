package cs_config

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	authSvc     iface.IAuthService
	csConfigSvc iface.ICsConfigService
}

type UpdateRemindParams struct {
	MaxMember           int64  `json:"max_member" binding:"required"`
	MemberPendingExpire int64  `json:"member_pending_expire" binding:"required"`
	GreetingText        string `json:"greeting_text" binding:"required"`
}

func (h *handler) UpdateCsConfig(c *gin.Context) {
	var (
		err           error
		requestParams UpdateRemindParams
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

	config := types.CsConfig{
		MaxMember:           requestParams.MaxMember,
		MemberPendingExpire: requestParams.MemberPendingExpire,
		GreetingText:        requestParams.GreetingText,
	}

	if err = h.csConfigSvc.UpdateCsConfig(ctx, staffInfo.ID, config); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetCsConfig(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	config, err := h.csConfigSvc.GetCsConfig(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, config)
}

func NewHandler(authSvc iface.IAuthService, csConfigSvc iface.ICsConfigService) *handler {
	return &handler{
		authSvc:     authSvc,
		csConfigSvc: csConfigSvc,
	}
}
