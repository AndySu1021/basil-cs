package lua

import (
	"context"
	"strconv"
	"time"
)

func (s *script) RemoveMemberToken(ctx context.Context, name string) error {
	return s.removeMemberToken.Run(ctx, s.redisClient.GetClient(), []string{name}).Err()
}

func (s *script) RemoveStaffToken(ctx context.Context, id int64, username string) error {
	return s.removeStaffToken.Run(ctx, s.redisClient.GetClient(), []string{strconv.FormatInt(id, 10), username}).Err()
}

func (s *script) SetToken(ctx context.Context, clientType string, name string, token string, value interface{}, duration time.Duration) error {
	expire := int64(duration / time.Second)
	return s.setToken.Run(ctx, s.redisClient.GetClient(), []string{clientType, name, token}, value, expire).Err()
}

func (s *script) StaffAssignRoom(ctx context.Context, id int64) error {
	return s.staffAssignRoom.Run(ctx, s.redisClient.GetClient(), []string{strconv.FormatInt(id, 10)}).Err()
}

func (s *script) StaffRegister(ctx context.Context, id int64) error {
	return s.staffRegister.Run(ctx, s.redisClient.GetClient(), []string{strconv.FormatInt(id, 10)}).Err()
}

func (s *script) StaffCloseRoom(ctx context.Context, staffId int64, memberName string) error {
	return s.staffCloseRoom.Run(ctx, s.redisClient.GetClient(), []string{strconv.FormatInt(staffId, 10), memberName}).Err()
}

func (s *script) StaffTransferRoom(ctx context.Context, staffId, toStaffId int64) error {
	return s.staffTransferRoom.Run(ctx, s.redisClient.GetClient(), []string{strconv.FormatInt(staffId, 10), strconv.FormatInt(toStaffId, 10)}).Err()
}
