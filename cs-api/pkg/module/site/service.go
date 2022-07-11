package site

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListSite(ctx context.Context) ([]model.ListSiteRow, error) {
	return s.repo.ListSite(ctx)
}

func (s *service) GetSite(ctx context.Context, siteId int64) (model.GetSiteRow, error) {
	return s.repo.GetSite(ctx, siteId)
}

func (s *service) CreateSite(ctx context.Context, params model.CreateSiteParams) error {
	return s.repo.CreateSite(ctx, params)
}

func (s *service) UpdateSite(ctx context.Context, params model.UpdateSiteParams) error {
	return s.repo.UpdateSite(ctx, params)
}

func (s *service) DeleteSite(ctx context.Context, siteId int64) error {
	return s.repo.DeleteSite(ctx, siteId)
}

func (s *service) GetSiteByCode(ctx context.Context, code string) (model.GetSiteByCodeRow, error) {
	return s.repo.GetSiteByCode(ctx, code)
}

func NewService(repo iface.IRepository) iface.ISiteService {
	return &service{
		repo: repo,
	}
}
