package notice

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

	routes.GET("/notices",
		p.AuthSvc.CheckPermission("ListNotice"),
		p.R.Op("ListNotice"),
		p.H.ListNotice,
	)

	routes.POST("/notice",
		p.AuthSvc.CheckPermission("ListNotice"),
		p.R.Op("CreateNotice"),
		p.H.CreateNotice,
	)

	routes.GET("/notice/:id",
		p.AuthSvc.CheckPermission("ListNotice"),
		p.R.Op("GetNotice"),
		p.H.GetNotice,
	)

	routes.PUT("/notice/:id",
		p.AuthSvc.CheckPermission("ListNotice"),
		p.R.Op("UpdateNotice"),
		p.H.UpdateNotice,
	)

	routes.DELETE("/notice/:id",
		p.AuthSvc.CheckPermission("ListNotice"),
		p.R.Op("DeleteNotice"),
		p.H.DeleteNotice,
	)

	p.Engine.GET("/api/notice/latest",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeMember),
		p.R.Op("GetLatestNotice"),
		p.H.GetLatestNotice,
	)
}
