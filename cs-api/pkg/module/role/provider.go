package role

import (
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		NewService,
		NewHandler,
	),
	fx.Invoke(
		InitTransport,
	),
)
