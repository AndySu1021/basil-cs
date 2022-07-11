package repository

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"database/sql"
	"go.uber.org/fx"
)

type repository struct {
	*model.Queries
	db *sql.DB
}

type Params struct {
	fx.In

	Queries *model.Queries
	DB      *sql.DB
}

func NewRepository(p Params) iface.IRepository {
	return &repository{
		Queries: p.Queries,
		db:      p.DB,
	}
}
