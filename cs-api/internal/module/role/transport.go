package role

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

	routes.GET("/roles",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("ListRole"),
		p.H.ListRole,
	)

	routes.GET("/roles/all",
		p.R.Op("GetAllRoles"),
		p.H.GetAllRoles,
	)

	routes.POST("/role",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("CreateRole"),
		p.H.CreateRole,
	)

	routes.GET("/role/:id",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("GetRole"),
		p.H.GetRole,
	)

	routes.PUT("/role/:id",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("UpdateRole"),
		p.H.UpdateRole,
	)

	routes.DELETE("/role/:id",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("DeleteRole"),
		p.H.DeleteRole,
	)
}
