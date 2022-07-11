package faq

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListFAQ(ctx context.Context, params types.ListFAQParams) ([]types.ListFAQRow, int64, error) {
	return s.repo.ListFAQ(ctx, params)
}

func (s *service) GetFAQ(ctx context.Context, faqId int64) (model.GetFAQRow, error) {
	return s.repo.GetFAQ(ctx, faqId)
}

func (s *service) CreateFAQ(ctx context.Context, params model.CreateFAQParams) error {
	return s.repo.CreateFAQ(ctx, params)
}

func (s *service) UpdateFAQ(ctx context.Context, params model.UpdateFAQParams) error {
	return s.repo.UpdateFAQ(ctx, params)
}

func (s *service) DeleteFAQ(ctx context.Context, faqId int64) error {
	return s.repo.DeleteFAQ(ctx, faqId)
}

func (s *service) ListAvailableFAQ(ctx context.Context) ([]model.ListAvailableFAQRow, error) {
	return s.repo.ListAvailableFAQ(ctx)
}

func NewService(Repo iface.IRepository) iface.IFAQService {
	return &service{
		repo: Repo,
	}
}
