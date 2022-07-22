package chat

import (
	"context"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"errors"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
)

type StaffDispatcher struct {
	redis     ifaceTool.IRedis
	lua       iface.ILuaScript
	staffs    map[int64]*StaffClient // staffID map to conn
	maxMember int64
	lock      *sync.RWMutex
}

func (sd *StaffDispatcher) getStaff(staffId int64) *StaffClient {
	if staffId == 0 {
		return nil
	}

	return sd.staffs[staffId]
}

func (sd *StaffDispatcher) register(sc *StaffClient) error {
	if err := sd.lua.StaffRegister(context.Background(), sc.ID); err != nil {
		return err
	}
	sd.staffs[sc.ID] = sc
	return nil
}

func (sd *StaffDispatcher) unregister(staffId int64) {
	delete(sd.staffs, staffId)
}

func (sd *StaffDispatcher) dispatch(staffId int64) (int64, error) {
	ctx := context.Background()
	// 如果用戶不小心斷線重新連回來時，優先找回原本的客服
	_, err := sd.redis.ZRank(ctx, types.RedisKeyStaffDispatch, strconv.FormatInt(staffId, 10)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}
	if !errors.Is(err, redis.Nil) {
		return staffId, nil
	}

	// 獲取當前服務數最少的客服
	result, err := sd.redis.ZRangeByScore(ctx, types.RedisKeyStaffDispatch, &redis.ZRangeBy{
		Min:    strconv.Itoa(0),
		Max:    strconv.FormatInt(sd.maxMember, 10),
		Offset: 0,
		Count:  1,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}

	return strconv.ParseInt(result[0], 10, 64)
}

func (sd *StaffDispatcher) setMaxMember(maxMember int64) {
	sd.maxMember = maxMember
}

func (sd *StaffDispatcher) assignRoom(staffId int64) error {
	return sd.lua.StaffAssignRoom(context.Background(), staffId)
}

func (sd *StaffDispatcher) removeRoom(staffId int64) error {
	return sd.redis.ZAddArgsIncr(context.Background(), types.RedisKeyStaffDispatch, redis.ZAddArgs{
		XX: true,
		Members: []redis.Z{{
			Score:  -1,
			Member: staffId,
		}},
	}).Err()
}

func NewStaffDispatcher(redis ifaceTool.IRedis, lua iface.ILuaScript, csConfigSvc iface.ICsConfigService) *StaffDispatcher {
	config, err := csConfigSvc.GetCsConfig(context.Background())
	if err != nil {
		return nil
	}

	return &StaffDispatcher{
		redis:     redis,
		lua:       lua,
		staffs:    make(map[int64]*StaffClient),
		maxMember: config.MaxMember,
		lock:      &sync.RWMutex{},
	}
}
