package chat

import (
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		NewNotifier,
		NewStaffDispatcher,
		NewClientManager,
		NewHandler,
		NewRedisWorker,
	),
	fx.Invoke(
		InitTransport,
		InitRedisSubscriber,
	),
)
