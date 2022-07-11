package common

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/module/tool"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

type Params struct {
	fx.In

	Engine  *gin.Engine
	AuthSvc iface.IAuthService
	H       *handler
	R       *tool.RequestInstrument
}

func InitTransport(p Params) {
	p.Engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	p.Engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	p.Engine.GET("/api/dashboard/panel-data",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("GetPanelData"),
		p.H.GetPanelData,
	)

	p.Engine.POST("/api/upload",
		p.R.Op("UploadFile"),
		p.H.UploadFile,
	)
}
