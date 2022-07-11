package report

import (
	"database/sql"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		NewService,
		NewHandler,
	),
	fx.Invoke(
		InitTransport,
		registerCron,
	),
)

func registerCron(db *sql.DB) {
	c := cron.New()

	// DailyStaffReport
	if _, err := c.AddJob("*/10 * * * *", DailyStaff{db}); err != nil {
		return
	}

	// DailyTagReport
	if _, err := c.AddJob("*/10 * * * *", DailyTag{db}); err != nil {
		return
	}

	c.Start()
}
