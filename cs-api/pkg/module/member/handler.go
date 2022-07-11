package member

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc   iface.IAuthService
	memberSvc iface.IMemberService
}

func (h *handler) ListMember(c *gin.Context) {
	var (
		err    error
		params types.ListMemberParams
		ctx    = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	reminds, total, err := h.memberSvc.ListMember(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params.Pagination.Total = total

	ginTool.SuccessWithPagination(c, reminds, params.Pagination)
}

type UpdateMemberInfoParams struct {
	RealName string `json:"real_name" binding:""`
	Email    string `json:"email" binding:""`
	Mobile   string `json:"mobile" binding:""`
	Note     string `json:"note" binding:""`
}

func (h *handler) UpdateMemberInfo(c *gin.Context) {
	var (
		err           error
		requestParams UpdateMemberInfoParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	params := model.UpdateMemberInfoParams{
		RealName:  requestParams.RealName,
		Email:     requestParams.Email,
		Mobile:    requestParams.Mobile,
		Note:      requestParams.Note,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.memberSvc.UpdateMemberInfo(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func NewHandler(authSvc iface.IAuthService, memberSvc iface.IMemberService) *handler {
	return &handler{
		authSvc:   authSvc,
		memberSvc: memberSvc,
	}
}
