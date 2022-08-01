package lua

import (
	iface "cs-api/internal/interface"
	_ "embed"
	iface2 "github.com/AndySu1021/go-util/interface"
	"github.com/go-redis/redis/v8"
)

//go:embed removeMemberToken.lua
var removeMemberTokenScript string

//go:embed removeStaffToken.lua
var removeStaffTokenScript string

//go:embed setToken.lua
var setTokenScript string

//go:embed staffAssignRoom.lua
var staffAssignRoomScript string

//go:embed staffRegister.lua
var staffRegisterScript string

//go:embed staffCloseRoom.lua
var staffCloseRoomScript string

//go:embed staffTransferRoom.lua
var staffTransferRoomScript string

type script struct {
	redisClient       iface2.IRedis
	removeMemberToken *redis.Script
	removeStaffToken  *redis.Script
	setToken          *redis.Script
	staffAssignRoom   *redis.Script
	staffRegister     *redis.Script
	staffCloseRoom    *redis.Script
	staffTransferRoom *redis.Script
}

func NewLua(redisClient iface2.IRedis) iface.ILuaScript {
	return &script{
		redisClient:       redisClient,
		removeMemberToken: redis.NewScript(removeMemberTokenScript),
		removeStaffToken:  redis.NewScript(removeStaffTokenScript),
		setToken:          redis.NewScript(setTokenScript),
		staffAssignRoom:   redis.NewScript(staffAssignRoomScript),
		staffRegister:     redis.NewScript(staffRegisterScript),
		staffCloseRoom:    redis.NewScript(staffCloseRoomScript),
		staffTransferRoom: redis.NewScript(staffTransferRoomScript),
	}
}
