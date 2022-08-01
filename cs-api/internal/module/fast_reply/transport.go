package fast_reply

import (
	iface "cs-api/internal/interface"
	"cs-api/internal/module/tool"
	"cs-api/internal"
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

	routes.GET("/fast-replies",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("ListFastReply"),
		p.H.ListFastReply,
	)

	routes.POST("/fast-reply",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("CreateFastReply"),
		p.H.CreateFastReply,
	)

	routes.GET("/fast-reply/:id",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("GetFastReply"),
		p.H.GetFastReply,
	)

	routes.PUT("/fast-reply/:id",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("UpdateFastReply"),
		p.H.UpdateFastReply,
	)

	routes.DELETE("/fast-reply/:id",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("DeleteFastReply"),
		p.H.DeleteFastReply,
	)

	routes.GET("/fast-reply/group",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("ListFastReplyGroup"),
		p.H.ListFastReplyGroup,
	)

	routes.GET("/fast-reply/categories",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("ListFastReplyCategory"),
		p.H.ListFastReplyCategory,
	)

	routes.POST("/fast-reply/category",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("CreateFastReplyCategory"),
		p.H.CreateFastReplyCategory,
	)
}
