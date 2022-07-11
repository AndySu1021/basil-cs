package message

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"errors"
)

type service struct {
	repo iface.IRepository
}

func (s *service) CreateMessage(ctx context.Context, params model.CreateMessageParams) error {
	return s.repo.CreateMessage(ctx, params)
}

func (s *service) ListRoomMessage(ctx context.Context, params interface{}) (messages []model.Message, err error) {
	messages = make([]model.Message, 0)

	switch data := params.(type) {
	case model.ListMemberRoomMessageParams:
		messages, err = s.repo.ListMemberRoomMessage(ctx, data)
	case model.ListStaffRoomMessageParams:
		messages, err = s.repo.ListStaffRoomMessage(ctx, data)
	default:
		return make([]model.Message, 0), errors.New("wrong client type")
	}

	if err != nil {
		return
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return
}

func (s *service) ListMessage(ctx context.Context, params types.ListMessageParams) ([]types.ListMessageRow, int64, error) {
	return s.repo.ListMessage(ctx, params)
}

func NewService(repo iface.IRepository) iface.IMessageService {
	return &service{
		repo: repo,
	}
}
