package site

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

	routes.GET("/sites",
		p.AuthSvc.CheckPermission("ListSite"),
		p.R.Op("ListSite"),
		p.H.ListSite,
	)

	routes.POST("/site",
		p.AuthSvc.CheckPermission("ListSite"),
		p.R.Op("CreateSite"),
		p.H.CreateSite,
	)

	routes.GET("/site/:id",
		p.AuthSvc.CheckPermission("ListSite"),
		p.R.Op("GetSite"),
		p.H.GetSite,
	)

	routes.PUT("/site/:id",
		p.AuthSvc.CheckPermission("ListSite"),
		p.R.Op("UpdateSite"),
		p.H.UpdateSite,
	)

	routes.DELETE("/site/:id",
		p.AuthSvc.CheckPermission("ListSite"),
		p.R.Op("DeleteSite"),
		p.H.DeleteSite,
	)
}
