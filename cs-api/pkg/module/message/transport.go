package message

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

	routes.GET("/messages",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("ListMessage"),
		p.H.ListMessage,
	)

	routes.GET("/room-messages/staff",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("ListStaffRoomMessage"),
		p.H.ListStaffRoomMessage,
	)

	routes.GET("/room-messages/member",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeMember),
		p.R.Op("ListMemberRoomMessage"),
		p.H.ListMemberRoomMessage,
	)
}
