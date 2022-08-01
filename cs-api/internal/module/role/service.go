package role

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"errors"
	"fmt"
	ifaceTool "github.com/AndySu1021/go-util/interface"
)

type service struct {
	repo  iface.IRepository
	redis ifaceTool.IRedis
}

func (s *service) ListRole(ctx context.Context, params types.ListRoleParams) ([]types.ListRoleRow, int64, error) {
	return s.repo.ListRole(ctx, params)
}

func (s *service) GetAllRoles(ctx context.Context) ([]model.GetAllRolesRow, error) {
	return s.repo.GetAllRoles(ctx)
}

func (s *service) GetRole(ctx context.Context, roleId int64) (model.Role, error) {
	return s.repo.GetRole(ctx, roleId)
}

func (s *service) CreateRole(ctx context.Context, params model.CreateRoleParams) error {
	return s.repo.CreateRole(ctx, params)
}

func (s *service) UpdateRole(ctx context.Context, params model.UpdateRoleParams) error {
	if err := s.repo.UpdateRole(ctx, params); err != nil {
		return err
	}
	return s.redis.Del(ctx, fmt.Sprintf("role:%d", params.ID)).Err()
}

func (s *service) DeleteRole(ctx context.Context, roleId int64) error {
	count, err := s.repo.GetStaffCountByRoleId(ctx, roleId)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("unable to delete")
	}

	return s.repo.DeleteRole(ctx, roleId)
}

func NewService(repo iface.IRepository, redis ifaceTool.IRedis) iface.IRoleService {
	return &service{
		repo:  repo,
		redis: redis,
	}
}
