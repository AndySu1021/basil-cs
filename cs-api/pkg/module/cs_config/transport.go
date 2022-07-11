package cs_config

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

	routes.GET("/cs-config",
		p.AuthSvc.CheckPermission("GetCsConfig"),
		p.R.Op("GetCsConfig"),
		p.H.GetCsConfig,
	)

	routes.PUT("/cs-config",
		p.AuthSvc.CheckPermission("GetCsConfig"),
		p.R.Op("UpdateCsConfig"),
		p.H.UpdateCsConfig,
	)
}
