package member

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
	routes := p.Engine.Group("/api", p.AuthSvc.SetClientInfo(internal.ClientTypeStaff))

	routes.GET("/members",
		p.AuthSvc.CheckPermission("ListMember"),
		p.R.Op("ListMember"),
		p.H.ListMember,
	)

	routes.PATCH("/member/:id/info",
		p.R.Op("UpdateMemberInfo"),
		p.H.UpdateMemberInfo,
	)
}
