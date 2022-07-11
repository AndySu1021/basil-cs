package faq

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
	routes := p.Engine.Group("/api", p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff))

	routes.GET("/faqs",
		p.AuthSvc.CheckPermission("ListFAQ"),
		p.R.Op("ListFAQ"),
		p.H.ListFAQ,
	)

	routes.POST("/faq",
		p.AuthSvc.CheckPermission("ListFAQ"),
		p.R.Op("CreateFAQ"),
		p.H.CreateFAQ,
	)

	routes.GET("/faq/:id",
		p.AuthSvc.CheckPermission("ListFAQ"),
		p.R.Op("GetFAQ"),
		p.H.GetFAQ,
	)

	routes.PUT("/faq/:id",
		p.AuthSvc.CheckPermission("ListFAQ"),
		p.R.Op("UpdateFAQ"),
		p.H.UpdateFAQ,
	)

	routes.DELETE("/faq/:id",
		p.AuthSvc.CheckPermission("ListFAQ"),
		p.R.Op("DeleteFAQ"),
		p.H.DeleteFAQ,
	)

	p.Engine.GET("/api/available-faqs",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeMember),
		p.R.Op("ListAvailableFAQ"),
		p.H.ListAvailableFAQ,
	)
}
