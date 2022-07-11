package room

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	ginTool "github.com/AndySu1021/go-util/gin"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"go.uber.org/fx"
	"time"
)

type service struct {
	redis ifaceTool.IRedis
	lua   iface.ILuaScript
	repo  iface.IRepository
}

func (s *service) CreateRoom(ctx context.Context, source types.RoomSource, userAgent string, memberId int64) (roomId int64, err error) {
	room, err := s.repo.GetMemberAvailableRoom(ctx, memberId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	if room.ID != 0 {
		roomId = room.ID
		return
	}

	ip := ctx.Value(ginTool.ContextKeyClientIp).(string)
	now := time.Now().UTC()
	result, err := s.repo.CreateRoom(ctx, model.CreateRoomParams{
		StaffID:   0,
		MemberID:  memberId,
		Ip:        ip,
		Source:    source,
		UserAgent: userAgent,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return
	}
	roomId, err = result.LastInsertId()
	if err != nil {
		return
	}

	return
}

func (s *service) AcceptRoom(ctx context.Context, staffId int64, roomId int64) error {
	var (
		err  error
		room model.GetRoomRow
	)

	room, err = s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.Status == types.RoomStatusClosed {
		return errors.New("wrong room status")
	}

	if err = s.repo.AcceptRoom(ctx, model.AcceptRoomParams{
		StaffID: staffId,
		ID:      roomId,
		AcceptedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}); err != nil {
		return err
	}

	// publish event to redis
	event := pkg.ClientEventInfo{
		Event: pkg.ClientEventAcceptRoom,
		Payload: pkg.ClientEventPayload{
			StaffID: &staffId,
			RoomID:  &roomId,
		},
	}
	payload, _ := json.Marshal(event)
	return s.redis.Publish(ctx, types.RedisKeyEventClient, payload)
}

func (s *service) CloseRoom(ctx context.Context, staffId int64, roomId int64, tagId int64) error {
	var (
		err  error
		room model.GetRoomRow
	)

	room, err = s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.StaffID != staffId {
		return errors.New("operation error")
	}

	if room.Status != types.RoomStatusServing {
		return errors.New("wrong room status")
	}

	if err = s.repo.CloseRoom(ctx, model.CloseRoomParams{
		TagID:    tagId,
		ClosedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		ID:       roomId,
	}); err != nil {
		return err
	}

	if err = s.lua.StaffCloseRoom(ctx, staffId, fmt.Sprintf("%d-%s", room.SiteID, room.MemberName)); err != nil {
		return err
	}

	event := pkg.ClientEventInfo{
		Event: pkg.ClientEventCloseRoom,
		Payload: pkg.ClientEventPayload{
			StaffID: &room.StaffID,
			RoomID:  &roomId,
		},
	}

	payload, _ := json.Marshal(event)
	return s.redis.Publish(ctx, types.RedisKeyEventClient, payload)
}

func (s *service) UpdateRoomScore(ctx context.Context, roomId int64, score int32) error {
	return s.repo.UpdateRoomScore(ctx, model.UpdateRoomScoreParams{
		Score: score,
		ID:    roomId,
	})
}

func (s *service) ListRoom(ctx context.Context, params types.ListRoomParams) ([]types.ListRoomRow, int64, error) {
	return s.repo.ListRoom(ctx, params)
}

func (s *service) ListStaffServingRoom(ctx context.Context, staffId int64) (result []types.StaffRoom, err error) {
	rooms, err := s.repo.ListStaffServingRoom(ctx, staffId)

	result = make([]types.StaffRoom, 0, len(rooms))
	for _, room := range rooms {
		tmp := types.StaffRoom{
			ID:           room.ID,
			Status:       room.Status,
			MemberID:     room.MemberID,
			MemberName:   room.MemberName,
			OnlineStatus: room.OnlineStatus,
		}
		lastMessage, err2 := s.repo.GetLastRoomMessage(ctx, room.ID)
		if err2 != nil && err2 != sql.ErrNoRows {
			return nil, err2
		}
		if err2 != sql.ErrNoRows {
			tmp.OpType = lastMessage.OpType
			tmp.ContentType = lastMessage.ContentType
			tmp.Content = lastMessage.Content
		}
		result = append(result, tmp)
	}

	return
}

func (s *service) ListQueuingRoom(ctx context.Context) ([]model.ListQueuingRoomRow, error) {
	return s.repo.ListQueuingRoom(ctx)
}

// GetStaffServingRoomCount 獲取客服尚未關閉服務的 房間ID 數量
func (s *service) GetStaffServingRoomCount(ctx context.Context, staffId int64) (int64, error) {
	return s.repo.GetStaffServingRoomCount(ctx, staffId)
}

func (s *service) TransferRoom(ctx context.Context, staffId, roomId, toStaffId int64) error {
	room, err := s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.StaffID != staffId {
		return errors.New("operation error")
	}

	if room.StaffID == toStaffId {
		return errors.New("no need to transfer")
	}

	staff, err := s.repo.GetStaff(ctx, toStaffId)
	if err != nil {
		return err
	}

	if staff.OnlineStatus != types.StaffOnlineStatusOnline {
		return errors.New("staff not available")
	}

	if err = s.repo.UpdateRoomStaff(ctx, model.UpdateRoomStaffParams{
		StaffID: toStaffId,
		ID:      roomId,
	}); err != nil {
		return err
	}

	if err = s.lua.StaffTransferRoom(ctx, staffId, toStaffId); err != nil {
		return err
	}

	// send event to redis
	event := pkg.ClientEventInfo{
		Event: pkg.ClientEventTransferRoom,
		Payload: pkg.ClientEventPayload{
			StaffID: &toStaffId,
			RoomID:  &roomId,
		},
	}
	payload, _ := json.Marshal(event)
	return s.redis.Publish(ctx, types.RedisKeyEventClient, payload)
}

func (s *service) GetMemberAvailableRoom(ctx context.Context, memberId int64) (model.Room, error) {
	return s.repo.GetMemberAvailableRoom(ctx, memberId)
}

func (s *service) GetRoomInfo(ctx context.Context, roomId int64) (model.GetRoomInfoRow, error) {
	return s.repo.GetRoomInfo(ctx, roomId)
}

func (s *service) UpdateRoomStatus(ctx context.Context, params model.UpdateRoomStatusParams) error {
	return s.repo.UpdateRoomStatus(ctx, params)
}

type ServiceParams struct {
	fx.In

	Redis     ifaceTool.IRedis
	Lua       iface.ILuaScript
	MemberSvc iface.IMemberService
	SiteSvc   iface.ISiteService
	Repo      iface.IRepository
}

func NewService(p ServiceParams) iface.IRoomService {
	return &service{
		redis: p.Redis,
		lua:   p.Lua,
		repo:  p.Repo,
	}
}
