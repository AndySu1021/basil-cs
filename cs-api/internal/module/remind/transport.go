package remind

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

	routes.GET("/reminds",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("ListRemind"),
		p.H.ListRemind,
	)

	routes.POST("/remind",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("CreateRemind"),
		p.H.CreateRemind,
	)

	routes.GET("/remind/:id",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("GetRemind"),
		p.H.GetRemind,
	)

	routes.PUT("/remind/:id",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("UpdateRemind"),
		p.H.UpdateRemind,
	)

	routes.DELETE("/remind/:id",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("DeleteRemind"),
		p.H.DeleteRemind,
	)

	routes.GET("/active-reminds",
		p.R.Op("ListActiveRemind"),
		p.H.ListActiveRemind,
	)
}
