package auth

import (
	"cs-api/internal"
	iface "cs-api/internal/interface"
	"cs-api/internal/module/tool"
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
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("GetStaffInfo"),
		p.H.GetStaffInfo,
	)

	routes.POST("/auth/logout",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("Logout"),
		p.H.Logout,
	)
}
