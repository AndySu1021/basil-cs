package report

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/module/tool"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Engine  *gin.Engine
	AuthSvc iface.IAuthService
	H       *handler
	R       *tool.RequestInstrument
}

func InitTransport(p Params) {
	routes := p.Engine.Group("/api", p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff))

	routes.GET("/report/daily/tag",
		p.AuthSvc.CheckPermission("ListDailyTagReport"),
		p.R.Op("ListDailyTagReport"),
		p.H.ListDailyTagReport,
	)

	routes.GET("/report/daily/staff",
		p.AuthSvc.CheckPermission("ListDailyStaffReport"),
		p.R.Op("ListDailyStaffReport"),
		p.H.ListDailyStaffReport,
	)
}
