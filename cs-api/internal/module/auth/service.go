package auth

import (
	"context"
	"crypto/md5"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/internal"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	errTool "github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/AndySu1021/go-util/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/rs/xid"
	"time"
)

type service struct {
	redis  ifaceTool.IRedis
	lua    iface.ILuaScript
	repo   iface.IRepository
	config *config.Config
}

func (s *service) Login(ctx context.Context, username, password string) (internal.ClientInfo, error) {
	staff, err := s.repo.StaffLogin(ctx, model.StaffLoginParams{
		Username: username,
		Password: helper.EncryptPassword(password, s.config.Salt),
	})
	if err != nil {
		return internal.ClientInfo{}, err
	}

	now := time.Now().UTC()
	if err = s.repo.UpdateStaffLogin(ctx, model.UpdateStaffLoginParams{
		ServingStatus: types.StaffServingStatusClosed,
		OnlineStatus:  types.StaffOnlineStatusOnline,
		LastLoginTime: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            staff.ID,
	}); err != nil {
		return internal.ClientInfo{}, err
	}

	token := genToken()
	var permissions []string
	if err = json.Unmarshal(staff.Permissions, &permissions); err != nil {
		return internal.ClientInfo{}, err
	}

	staffInfo := internal.ClientInfo{
		ID:          staff.ID,
		Type:        internal.ClientTypeStaff,
		Name:        staff.Name,
		Username:    staff.Username,
		RoleID:      staff.RoleID,
		Permissions: permissions,
		Token:       token,
	}

	result, err := json.Marshal(staffInfo)
	if err != nil {
		return internal.ClientInfo{}, err
	}

	err = s.lua.RemoveStaffToken(ctx, staff.ID, staff.Username)
	if err != nil {
		logger.Logger.Errorf("clear token error: %s\n", err)
		return internal.ClientInfo{}, err
	}

	err = s.lua.SetToken(ctx, "staff", staff.Username, token, result, 24*time.Hour)
	if err != nil {
		logger.Logger.Errorf("set token error: %s\n", err)
		return internal.ClientInfo{}, err
	}

	return staffInfo, nil
}

func (s *service) GetStaffInfo(ctx context.Context, staffId int64) (model.GetStaffRow, error) {
	return s.repo.GetStaff(ctx, staffId)
}

func (s *service) Logout(ctx context.Context, staffInfo internal.ClientInfo) error {
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

func (s *service) SetClientInfo(clientType internal.ClientType) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxKey := fmt.Sprintf("%s_info", clientType)
		token := c.GetHeader("X-Token")
		if token == "" {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		redisKey := getRedisKey(token, clientType)
		result, err := s.redis.Get(c.Request.Context(), redisKey).Result()
		if err != nil {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		var tmp internal.ClientInfo
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

func (s *service) GetClientInfo(ctx context.Context, clientType internal.ClientType) (internal.ClientInfo, error) {
	ctxKey := fmt.Sprintf("%s_info", clientType)
	clientInfo, ok := ctx.Value(ctxKey).(internal.ClientInfo)
	if clientInfo.ID == 0 || !ok {
		return internal.ClientInfo{}, errTool.ErrorAuth
	}

	return clientInfo, nil
}

func (s *service) CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		info, err := s.GetClientInfo(ctx, internal.ClientTypeStaff)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		roleKey := fmt.Sprintf("role:%d", info.RoleID)
		result, err := s.redis.Get(ctx, roleKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		var role model.Role
		if !errors.Is(err, redis.Nil) {
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
			if err = s.redis.SetEX(ctx, roleKey, roleBytes, 24*time.Hour).Err(); err != nil {
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

func getRedisKey(token string, clientType internal.ClientType) string {
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
