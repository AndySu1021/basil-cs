package tag

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

	routes.GET("/tags",
		p.AuthSvc.CheckPermission("ListTag"),
		p.R.Op("ListTag"),
		p.H.ListTag,
	)

	routes.POST("/tag",
		p.AuthSvc.CheckPermission("ListTag"),
		p.R.Op("CreateTag"),
		p.H.CreateTag,
	)

	routes.GET("/tag/:id",
		p.AuthSvc.CheckPermission("ListTag"),
		p.R.Op("GetTag"),
		p.H.GetTag,
	)

	routes.PUT("/tag/:id",
		p.AuthSvc.CheckPermission("ListTag"),
		p.R.Op("UpdateTag"),
		p.H.UpdateTag,
	)

	routes.DELETE("/tag/:id",
		p.AuthSvc.CheckPermission("ListTag"),
		p.R.Op("DeleteTag"),
		p.H.DeleteTag,
	)

	routes.GET("/available-tags",
		p.R.Op("ListAvailableTag"),
		p.H.ListAvailableTag,
	)
}
