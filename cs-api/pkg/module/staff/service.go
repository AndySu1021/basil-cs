package staff

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"errors"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/go-redis/redis/v8"
	"time"
)

type service struct {
	redis ifaceTool.IRedis
	repo  iface.IRepository
}

func (s *service) ListStaff(ctx context.Context, params types.ListStaffParams) ([]types.ListStaffRow, int64, error) {
	return s.repo.ListStaff(ctx, params)
}

func (s *service) GetStaff(ctx context.Context, staffId int64) (staff model.GetStaffRow, err error) {
	return s.repo.GetStaff(ctx, staffId)
}

func (s *service) CreateStaff(ctx context.Context, params model.CreateStaffParams) error {
	return s.repo.CreateStaff(ctx, params)
}

func (s *service) UpdateStaff(ctx context.Context, params interface{}) error {
	switch data := params.(type) {
	case model.UpdateStaffParams:
		return s.repo.UpdateStaff(ctx, data)
	case model.UpdateStaffWithPasswordParams:
		return s.repo.UpdateStaffWithPassword(ctx, data)
	case model.UpdateStaffAvatarParams:
		return s.repo.UpdateStaffAvatar(ctx, data)
	}
	return errors.New("params type error")
}

func (s *service) DeleteStaff(ctx context.Context, staffId int64) error {
	return s.repo.DeleteStaff(ctx, staffId)
}

func (s *service) UpdateStaffServingStatus(ctx context.Context, staffId int64, status types.StaffServingStatus) error {
	now := time.Now().UTC()

	params := model.UpdateStaffServingStatusParams{
		ServingStatus: status,
		UpdatedBy:     staffId,
		UpdatedAt:     now,
		ID:            staffId,
	}

	if err := s.repo.UpdateStaffServingStatus(ctx, params); err != nil {
		return err
	}

	switch status {
	case types.StaffServingStatusServing:
		roomCount, err := s.repo.GetStaffServingRoomCount(ctx, staffId)
		if err != nil {
			return err
		}
		return s.redis.ZAdd(ctx, types.RedisKeyStaffDispatch, &redis.Z{
			Score:  float64(roomCount),
			Member: staffId,
		})
	case types.StaffServingStatusClosed:
		return s.redis.ZAdd(ctx, types.RedisKeyStaffDispatch, &redis.Z{
			Score:  -1,
			Member: staffId,
		})
	}

	return nil
}

func (s *service) ListAvailableStaff(ctx context.Context, staffId int64) ([]model.ListAvailableStaffRow, error) {
	return s.repo.ListAvailableStaff(ctx, staffId)
}

func (s *service) GetAllStaffs(ctx context.Context) ([]model.GetAllStaffsRow, error) {
	return s.repo.GetAllStaffs(ctx)
}

func NewService(redis ifaceTool.IRedis, repo iface.IRepository) iface.IStaffService {
	return &service{
		redis: redis,
		repo:  repo,
	}
}
