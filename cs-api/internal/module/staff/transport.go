package staff

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

	routes.GET("/staffs",
		p.AuthSvc.CheckPermission("ListStaff"),
		p.R.Op("ListStaff"),
		p.H.ListStaff,
	)

	routes.POST("/staff",
		p.AuthSvc.CheckPermission("ListStaff"),
		p.R.Op("CreateStaff"),
		p.H.CreateStaff,
	)

	routes.GET("/staff/:id",
		p.AuthSvc.CheckPermission("ListStaff"),
		p.R.Op("GetStaff"),
		p.H.GetStaff,
	)

	routes.PUT("/staff/:id",
		p.AuthSvc.CheckPermission("ListStaff"),
		p.R.Op("UpdateStaff"),
		p.H.UpdateStaff,
	)

	routes.DELETE("/staff/:id",
		p.AuthSvc.CheckPermission("ListStaff"),
		p.R.Op("DeleteStaff"),
		p.H.DeleteStaff,
	)

	routes.PATCH("/staff/serving-status",
		p.R.Op("UpdateStaffServingStatus"),
		p.H.UpdateStaffServingStatus,
	)

	routes.PATCH("/staff/avatar",
		p.R.Op("UpdateStaffAvatar"),
		p.H.UpdateStaffAvatar,
	)

	routes.GET("/available-staffs",
		p.R.Op("ListAvailableStaff"),
		p.H.ListAvailableStaff,
	)

	routes.GET("/staffs/all",
		p.R.Op("GetAllStaffs"),
		p.H.GetAllStaffs,
	)
}
