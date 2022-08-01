package room

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

	routes.GET("/rooms",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.AuthSvc.CheckPermission("ListRoom"),
		p.R.Op("ListRoom"),
		p.H.ListRoom,
	)

	routes.GET("/staff-rooms",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("ListStaffServingRoom"),
		p.H.ListStaffServingRoom,
	)

	routes.GET("/queuing-rooms",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("ListQueuingRoom"),
		p.H.ListQueuingRoom,
	)

	routes.PATCH("/room/accept/:id",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("AcceptRoom"),
		p.H.AcceptRoom,
	)

	routes.PATCH("/room/close/:id",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("CloseRoom"),
		p.H.CloseRoom,
	)

	routes.PATCH("/room/transfer/:id",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("TransferRoom"),
		p.H.TransferRoom,
	)

	// 獲取房間詳細資訊
	routes.GET("/room/:id/info",
		p.AuthSvc.SetClientInfo(internal.ClientTypeStaff),
		p.R.Op("GetRoomInfo"),
		p.H.GetRoomInfo,
	)

	routes.PATCH("/room/score",
		p.AuthSvc.SetClientInfo(internal.ClientTypeMember),
		p.R.Op("UpdateRoomScore"),
		p.H.UpdateRoomScore,
	)

	routes.POST("/room",
		p.R.Op("CreateRoom"),
		p.H.CreateRoom,
	)
}
