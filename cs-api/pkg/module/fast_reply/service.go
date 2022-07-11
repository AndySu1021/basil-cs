package fast_reply

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListFastReply(ctx context.Context, params types.ListFastReplyParams) ([]types.ListFastReplyRow, int64, error) {
	return s.repo.ListFastReply(ctx, params)
}

func (s *service) GetFastReply(ctx context.Context, id int64) (model.FastReply, error) {
	return s.repo.GetFastReply(ctx, id)
}

func (s *service) CreateFastReply(ctx context.Context, params model.CreateFastReplyParams) error {
	return s.repo.CreateFastReply(ctx, params)
}

func (s *service) UpdateFastReply(ctx context.Context, params model.UpdateFastReplyParams) error {
	return s.repo.UpdateFastReply(ctx, params)
}

func (s *service) DeleteFastReply(ctx context.Context, id int64) error {
	return s.repo.DeleteFastReply(ctx, id)
}

func (s *service) ListCategory(ctx context.Context) ([]model.Constant, error) {
	return s.repo.ListFastReplyCategory(ctx)
}

func (s *service) CreateCategory(ctx context.Context, params model.CreateFastReplyCategoryParams) error {
	return s.repo.CreateFastReplyCategory(ctx, params)
}

func (s *service) ListFastReplyGroup(ctx context.Context) ([]pkg.FastReplyGroupItem, error) {
	messages, err := s.repo.GetAllAvailableFastReply(ctx)
	if err != nil {
		return nil, err
	}

	categoryMap := map[int64]pkg.FastReplyGroupItem{}
	for _, message := range messages {
		if _, ok := categoryMap[message.CategoryID]; !ok {
			tmp := pkg.FastReplyGroupItem{
				Category: pkg.FastReplyCategory{
					ID:   message.CategoryID,
					Name: message.Category,
				},
			}
			tmp.Items = append(tmp.Items, message)
			categoryMap[message.CategoryID] = tmp
		} else {
			tmp := categoryMap[message.CategoryID]
			tmp.Items = append(tmp.Items, message)
			categoryMap[message.CategoryID] = tmp
		}
	}

	group := make([]pkg.FastReplyGroupItem, 0)
	for _, v := range categoryMap {
		group = append(group, v)
	}

	return group, nil
}

func (s *service) CheckCategory(ctx context.Context, id int64) (interface{}, error) {
	return s.repo.CheckFastReplyCategory(ctx, id)
}

func NewService(Repo iface.IRepository) iface.IFastReplyService {
	return &service{
		repo: Repo,
	}
}
