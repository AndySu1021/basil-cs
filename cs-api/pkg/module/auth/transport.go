package auth

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
	routes := p.Engine.Group("/api")

	routes.POST("/auth/login",
		p.R.Op("Login"),
		p.H.Login,
	)

	routes.GET("/auth/staff-info",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("GetStaffInfo"),
		p.H.GetStaffInfo,
	)

	routes.POST("/auth/logout",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("Logout"),
		p.H.Logout,
	)
}
