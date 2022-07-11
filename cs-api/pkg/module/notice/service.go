package notice

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListNotice(ctx context.Context, params types.ListNoticeParams) ([]types.ListNoticeRow, int64, error) {
	return s.repo.ListNotice(ctx, params)
}

func (s *service) GetNotice(ctx context.Context, noticeId int64) (model.Notice, error) {
	return s.repo.GetNotice(ctx, noticeId)
}

func (s *service) CreateNotice(ctx context.Context, params model.CreateNoticeParams) error {
	return s.repo.CreateNotice(ctx, params)
}

func (s *service) UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) error {
	return s.repo.UpdateNotice(ctx, params)
}

func (s *service) DeleteNotice(ctx context.Context, noticeId int64) error {
	return s.repo.DeleteNotice(ctx, noticeId)
}

func (s *service) GetLatestNotice(ctx context.Context) (model.GetLatestNoticeRow, error) {
	return s.repo.GetLatestNotice(ctx)
}

func NewService(Repo iface.IRepository) iface.INoticeService {
	return &service{
		repo: Repo,
	}
}
