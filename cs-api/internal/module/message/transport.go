package message

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

	routes.GET("/messages",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("ListMessage"),
		p.H.ListMessage,
	)

	routes.GET("/room-messages/staff",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("ListStaffRoomMessage"),
		p.H.ListStaffRoomMessage,
	)

	routes.GET("/room-messages/member",
		p.AuthSvc.SetClientInfo(internal.ClientTypeMember),
		p.R.Op("ListMemberRoomMessage"),
		p.H.ListMemberRoomMessage,
	)
}
