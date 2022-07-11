package auth

import (
	"context"
	"crypto/md5"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/AndySu1021/go-util/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"time"
)

type service struct {
	redis  ifaceTool.IRedis
	lua    iface.ILuaScript
	repo   iface.IRepository
	config *config.Config
}

func (s *service) Login(ctx context.Context, username, password string) (pkg.ClientInfo, error) {
	staff, err := s.repo.StaffLogin(ctx, model.StaffLoginParams{
		Username: username,
		Password: helper.EncryptPassword(password, s.config.Salt),
	})
	if err != nil {
		return pkg.ClientInfo{}, err
	}

	now := time.Now().UTC()
	if err = s.repo.UpdateStaffLogin(ctx, model.UpdateStaffLoginParams{
		ServingStatus: types.StaffServingStatusClosed,
		OnlineStatus:  types.StaffOnlineStatusOnline,
		LastLoginTime: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            staff.ID,
	}); err != nil {
		return pkg.ClientInfo{}, err
	}

	token := genToken()
	var permissions []string
	if err = json.Unmarshal(staff.Permissions, &permissions); err != nil {
		return pkg.ClientInfo{}, err
	}

	staffInfo := pkg.ClientInfo{
		ID:          staff.ID,
		Type:        pkg.ClientTypeStaff,
		Name:        staff.Name,
		Username:    staff.Username,
		RoleID:      staff.RoleID,
		Permissions: permissions,
		Token:       token,
	}

	result, err := json.Marshal(staffInfo)
	if err != nil {
		return pkg.ClientInfo{}, err
	}

	err = s.lua.RemoveStaffToken(ctx, staff.ID, staff.Username)
	if err != nil {
		logger.Logger.Errorf("clear token error: %s\n", err)
		return pkg.ClientInfo{}, err
	}

	err = s.lua.SetToken(ctx, "staff", staff.Username, token, result, 24*time.Hour)
	if err != nil {
		logger.Logger.Errorf("set token error: %s\n", err)
		return pkg.ClientInfo{}, err
	}

	return staffInfo, nil
}

func (s *service) GetStaffInfo(ctx context.Context, staffId int64) (model.GetStaffRow, error) {
	return s.repo.GetStaff(ctx, staffId)
}

func (s *service) Logout(ctx context.Context, staffInfo pkg.ClientInfo) error {
	params := model.UpdateStaffLogoutParams{
		ServingStatus: types.StaffServingStatusClosed,
		OnlineStatus:  types.StaffOnlineStatusOffline,
		UpdatedAt:     time.Now().UTC(),
		ID:            staffInfo.ID,
	}
	if err := s.repo.UpdateStaffLogout(ctx, params); err != nil {
		return err
	}

	if err := s.lua.RemoveStaffToken(ctx, staffInfo.ID, staffInfo.Name); err != nil {
		return err
	}

	return nil
}

func (s *service) SetClientInfo(clientType pkg.ClientType) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxKey := fmt.Sprintf("%s_info", clientType)
		token := c.GetHeader("X-Token")
		if token == "" {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		redisKey := getRedisKey(token, clientType)
		result, err := s.redis.Get(c.Request.Context(), redisKey)
		if err != nil || result == "" {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		var tmp pkg.ClientInfo
		err = json.Unmarshal([]byte(result), &tmp)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), ctxKey, tmp)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *service) GetClientInfo(ctx context.Context, clientType pkg.ClientType) (pkg.ClientInfo, error) {
	ctxKey := fmt.Sprintf("%s_info", clientType)
	clientInfo, ok := ctx.Value(ctxKey).(pkg.ClientInfo)
	if clientInfo.ID == 0 || !ok {
		return pkg.ClientInfo{}, errors.ErrorAuth
	}

	return clientInfo, nil
}

func (s *service) CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		info, err := s.GetClientInfo(ctx, pkg.ClientTypeStaff)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		roleKey := fmt.Sprintf("role:%d", info.RoleID)
		result, err := s.redis.Get(ctx, roleKey)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		var role model.Role
		if result != "" {
			if err = json.Unmarshal([]byte(result), &role); err != nil {
				ginTool.Error(c, err)
				c.Abort()
				return
			}
		} else {
			role, err = s.repo.GetRole(ctx, info.RoleID)
			if err != nil {
				ginTool.Error(c, err)
				c.Abort()
				return
			}
			roleBytes, err := json.Marshal(role)
			if err != nil {
				ginTool.Error(c, err)
				c.Abort()
				return
			}
			if err = s.redis.SetEX(ctx, roleKey, roleBytes, 24*time.Hour); err != nil {
				ginTool.Error(c, err)
				c.Abort()
				return
			}
		}

		if role.ID == 1 {
			c.Next()
			return
		}

		var permissions []string
		if err = json.Unmarshal(role.Permissions, &permissions); err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		for i := 0; i < len(permissions); i++ {
			if permission == permissions[i] {
				c.Next()
				return
			}
		}

		ginTool.ErrorPerm(c)
		c.Abort()
	}
}

func getRedisKey(token string, clientType pkg.ClientType) string {
	return fmt.Sprintf("token:%s:%s", clientType, token)
}

func genToken() string {
	str := time.Now().String() + xid.New().String()
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return str
}

func NewService(redis ifaceTool.IRedis, lua iface.ILuaScript, repo iface.IRepository, config *config.Config) iface.IAuthService {
	return &service{
		redis:  redis,
		lua:    lua,
		repo:   repo,
		config: config,
	}
}
