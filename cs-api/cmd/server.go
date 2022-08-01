package cmd

import (
	"context"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/internal/lua"
	"cs-api/internal/module"
	"cs-api/internal/repository"
	"database/sql"
	"errors"
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	"github.com/AndySu1021/go-util/logger"
	"github.com/AndySu1021/go-util/redis"
	"github.com/AndySu1021/go-util/storage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ServerCmd = &cobra.Command{
	Run: runServer,
	Use: "server",
}

func runServer(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	options := fx.Options(
		fx.Provide(
			config.NewConfig,
			gin.NewGin,
			lua.NewLua,
			db.NewDatabase,
			NewDBTX,
			model.New,
			repository.NewRepository,
			storage.NewStorage,
			NewContext,
		),
		fx.Invoke(
			logger.InitZapLogger,
			Migrate,
			Seed,
		),
	)

	app := fx.New(
		fx.NopLogger,
		options,
		module.Options,
		redis.Options,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		logger.Logger.Errorf("app start error: %s", err)
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	logger.Logger.Info("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		logger.Logger.Error("app stop err")
	}

	os.Exit(exitCode)
}

func NewDBTX(db *sql.DB) model.DBTX {
	return db
}

func Migrate(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

func Seed(cfg *config.Config, q *model.Queries) (err error) {
	ctx := context.Background()
	if err = q.TagSeeder(ctx); err != nil {
		return
	}
	if err = q.RoleSeeder(ctx); err != nil {
		return
	}
	if err = q.StaffSeeder(ctx, helper.EncryptPassword(cfg.AdminPassword, cfg.Salt)); err != nil {
		return
	}
	if err = q.ConstantSeeder(ctx); err != nil {
		return
	}
	return
}

func NewContext() context.Context {
	return context.Background()
}
