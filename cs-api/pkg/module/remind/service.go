package remind

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListRemind(ctx context.Context, params types.ListRemindParams) ([]types.ListRemindRow, int64, error) {
	return s.repo.ListRemind(ctx, params)
}

func (s *service) GetRemind(ctx context.Context, remindId int64) (model.Remind, error) {
	return s.repo.GetRemind(ctx, remindId)
}

func (s *service) CreateRemind(ctx context.Context, params model.CreateRemindParams) error {
	return s.repo.CreateRemind(ctx, params)
}

func (s *service) UpdateRemind(ctx context.Context, params model.UpdateRemindParams) error {
	return s.repo.UpdateRemind(ctx, params)
}

func (s *service) DeleteRemind(ctx context.Context, remindId int64) error {
	return s.repo.DeleteRemind(ctx, remindId)
}

func (s *service) ListActiveRemind(ctx context.Context) ([]model.ListActiveRemindRow, error) {
	return s.repo.ListActiveRemind(ctx)
}

func NewService(Repo iface.IRepository) iface.IRemindService {
	return &service{
		repo: Repo,
	}
}
