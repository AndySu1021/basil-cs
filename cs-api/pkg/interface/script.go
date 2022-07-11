package iface

import (
	"context"
	"time"
)

type ILuaScript interface {
	RemoveMemberToken(ctx context.Context, name string) error
	RemoveStaffToken(ctx context.Context, id int64, username string) error
	SetToken(ctx context.Context, clientType string, name string, token string, value interface{}, duration time.Duration) error
	StaffAssignRoom(ctx context.Context, id int64) error
	StaffRegister(ctx context.Context, id int64) error
	StaffCloseRoom(ctx context.Context, staffId int64, memberName string) error
	StaffTransferRoom(ctx context.Context, staffId, toStaffId int64) error
}
