package tag

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListTag(ctx context.Context, params types.ListTagParams) (tags []types.ListTagRow, total int64, err error) {
	return s.repo.ListTag(ctx, params)
}

func (s *service) GetTag(ctx context.Context, tagId int64) (tag model.GetTagRow, err error) {
	return s.repo.GetTag(ctx, tagId)
}

func (s *service) CreateTag(ctx context.Context, params model.CreateTagParams) error {
	return s.repo.CreateTag(ctx, params)
}

func (s *service) UpdateTag(ctx context.Context, params model.UpdateTagParams) error {
	return s.repo.UpdateTag(ctx, params)
}

func (s *service) DeleteTag(ctx context.Context, tagId int64) error {
	return s.repo.DeleteTag(ctx, tagId)
}

func (s *service) ListAvailableTag(ctx context.Context) ([]model.ListAvailableTagRow, error) {
	return s.repo.ListAvailableTag(ctx)
}

func NewService(repo iface.IRepository) iface.ITagService {
	return &service{
		repo: repo,
	}
}
