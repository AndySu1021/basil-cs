package report

import (
	iface "cs-api/internal/interface"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"time"
)

type handler struct {
	authSvc   iface.IAuthService
	reportSvc iface.IReportService
}

type ListDailyTagReportParams struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}

func (h *handler) ListDailyTagReport(c *gin.Context) {
	var (
		err           error
		requestParams ListDailyTagReportParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	startDate, endDate, err := parseDateRange(requestParams.StartDate, requestParams.EndDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	column, data, err := h.reportSvc.ListDailyTagReport(ctx, startDate, endDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{
		"columns": column,
		"items":   data,
	})
}

type ListDailyStaffReportParams struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}

func (h *handler) ListDailyStaffReport(c *gin.Context) {
	var (
		err           error
		requestParams ListDailyStaffReportParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	startDate, endDate, err := parseDateRange(requestParams.StartDate, requestParams.EndDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	data, err := h.reportSvc.ListDailyStaffReport(ctx, startDate, endDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, data)
}

func parseDateRange(startDateStr, endDateStr string) (startDate, endDate time.Time, err error) {
	startDate, err = time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return
	}

	endDate, err = time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return
	}

	return
}

func NewHandler(authSvc iface.IAuthService, reportSvc iface.IReportService) *handler {
	return &handler{
		authSvc:   authSvc,
		reportSvc: reportSvc,
	}
}
