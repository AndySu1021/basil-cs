package common

import (
	"context"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
)

type service struct {
	repo iface.IRepository
}

func (s *service) GetPanelData(ctx context.Context) (data types.PanelData, err error) {
	data = types.PanelData{
		StaffCount:   0,
		ServingCount: 0,
		QueuingCount: 0,
	}

	data.StaffCount, err = s.repo.GetOnlineStaffCount(ctx)
	if err != nil {
		return
	}

	data.ServingCount, err = s.repo.GetServingRoomCount(ctx)
	if err != nil {
		return
	}

	data.QueuingCount, err = s.repo.GetQueuingRoomCount(ctx)
	if err != nil {
		return
	}

	return
}

func NewService(Repo iface.IRepository) iface.ICommonService {
	return &service{
		repo: Repo,
	}
}
