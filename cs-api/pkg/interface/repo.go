package iface

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg/types"
	"database/sql"
)

type Callback func(ctx context.Context, tx *sql.Tx) error

type IRepository interface {
	model.Querier
	WithTx(tx *sql.Tx) model.Querier
	Transaction(ctx context.Context, f Callback) error
	ListTag(ctx context.Context, params types.ListTagParams) ([]types.ListTagRow, int64, error)
	ListStaff(ctx context.Context, params types.ListStaffParams) ([]types.ListStaffRow, int64, error)
	ListRoom(ctx context.Context, params types.ListRoomParams) ([]types.ListRoomRow, int64, error)
	ListRole(ctx context.Context, params types.ListRoleParams) ([]types.ListRoleRow, int64, error)
	ListRemind(ctx context.Context, params types.ListRemindParams) ([]types.ListRemindRow, int64, error)
	ListNotice(ctx context.Context, params types.ListNoticeParams) ([]types.ListNoticeRow, int64, error)
	ListMember(ctx context.Context, params types.ListMemberParams) ([]types.ListMemberRow, int64, error)
	ListFastReply(ctx context.Context, params types.ListFastReplyParams) ([]types.ListFastReplyRow, int64, error)
	ListFAQ(ctx context.Context, params types.ListFAQParams) ([]types.ListFAQRow, int64, error)
	ListMessage(ctx context.Context, params types.ListMessageParams) ([]types.ListMessageRow, int64, error)
}
