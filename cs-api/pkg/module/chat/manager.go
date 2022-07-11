package chat

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"fmt"
	"github.com/AndySu1021/go-util/logger"
	"go.uber.org/fx"
	"sync"
	"time"
)

type ClientManager struct {
	lock          *sync.RWMutex // 需要加鎖防止等待用戶超出上限
	memberClients *sync.Map     // roomId maps member client
	dispatcher    *StaffDispatcher
	notifier      *Notifier
	roomSvc       iface.IRoomService
	memberSvc     iface.IMemberService
	csConfig      types.CsConfig
	lua           iface.ILuaScript
	memberPool    *sync.Pool
	staffPool     *sync.Pool
}

func (w *ClientManager) Register(clientInfo pkg.ClientInfo) {
	switch clientInfo.Type {
	case pkg.ClientTypeStaff:
		staff := w.staffPool.Get().(*StaffClient)
		if err := staff.Reset(clientInfo, w); err != nil {
			logger.Logger.Errorf("reset staff error: %s", err)
			return
		}
		if err := w.dispatcher.register(staff); err != nil {
			logger.Logger.Errorf("register staff error: %s", err)
			return
		}
		go staff.SocketRead()
		go staff.SocketWrite()
	case pkg.ClientTypeMember:
		// prevent duplicate member client connection
		if old := w.GetMember(clientInfo.RoomID); old != nil {
			if err := w.Unregister(old); err != nil {
				logger.Logger.Errorf("unregister member error: %s", err)
				return
			}
		}

		member := w.memberPool.Get().(*MemberClient)
		if err := member.Reset(clientInfo, w); err != nil {
			logger.Logger.Errorf("reset member error: %s", err)
			return
		}

		if err := w.memberSvc.UpdateOnlineStatus(context.Background(), model.UpdateOnlineStatusParams{
			OnlineStatus: types.MemberOnlineStatusOnline,
			ID:           member.ID,
		}); err != nil {
			logger.Logger.Errorf("update member online status error: %s", err)
			return
		}

		w.JoinRoom(member)

		// 檢查用戶是否超過閒置時間
		go checkPendingTimeout(member)
		go member.SocketRead()
		go member.SocketWrite()

		// 發送問候語
		w.notifier.Greeting(w.csConfig.GreetingText, member)

		// 指派一位客服給用戶
		staffId, err := w.dispatcher.dispatch(member.StaffID)
		if err != nil {
			logger.Logger.Errorf("dispatch staff error: %s", err.Error())
			return
		}
		if staffId == 0 {
			if err = w.roomSvc.UpdateRoomStatus(context.Background(), model.UpdateRoomStatusParams{
				Status: types.RoomStatusQueuing,
				ID:     member.RoomID,
			}); err != nil {
				logger.Logger.Errorf("update room status error: %s", err.Error())
				return
			}
			w.notifier.NoStaff(member)
		} else {
			if member.StaffID != staffId {
				member.StaffID = staffId
				if err = w.roomSvc.AcceptRoom(context.Background(), staffId, member.RoomID); err != nil {
					logger.Logger.Errorf("accept room error: %s", err.Error())
					return
				}
				w.notifier.MemberJoin(staffId, member)
			}
		}
	}
}

func (w *ClientManager) Unregister(client Client) error {
	if client.GetStatus() == ClientStatusClosed {
		return nil
	}
	if client.GetType() == pkg.ClientTypeStaff {
		staff := client.(*StaffClient)
		staff.Status = ClientStatusClosed
		close(staff.SendChan)
		if err := staff.Socket.Close(); err != nil {
			logger.Logger.Errorf("close socket error: %s\n", err)
			return err
		}
		w.dispatcher.unregister(staff.ID)
		w.staffPool.Put(staff)
	} else if client.GetType() == pkg.ClientTypeMember {
		member := client.(*MemberClient)
		member.Status = ClientStatusClosed
		close(member.SendChan)
		close(member.Sending)
		if err := member.Socket.Close(); err != nil {
			logger.Logger.Errorf("handle unregister error: %s", err)
			return err
		}
		if err := w.memberSvc.UpdateOnlineStatus(context.Background(), model.UpdateOnlineStatusParams{
			OnlineStatus: types.MemberOnlineStatusOffline,
			ID:           member.ID,
		}); err != nil {
			logger.Logger.Errorf("update member online status error: %s", err)
			return err
		}
		w.memberClients.Delete(member.RoomID)
		w.memberPool.Put(member)
	}

	return nil
}

func (w *ClientManager) GetMember(roomId int64) *MemberClient {
	if member, ok := w.memberClients.Load(roomId); ok {
		return member.(*MemberClient)
	}
	return nil
}

func (w *ClientManager) JoinRoom(member *MemberClient) {
	w.memberClients.Store(member.RoomID, member)
}

func (w *ClientManager) CloseRoom(roomId int64) {
	if client, ok := w.memberClients.Load(roomId); ok {
		if err := w.Unregister(client.(*MemberClient)); err != nil {
			logger.Logger.Errorf("unregister staff client error: %s", err)
			return
		}
	}
}

type ClientManagerParams struct {
	fx.In

	RoomSvc     iface.IRoomService
	MemberSvc   iface.IMemberService
	CsConfigSvc iface.ICsConfigService
	Dispatcher  *StaffDispatcher
	Notifier    *Notifier
	Lua         iface.ILuaScript
}

func NewClientManager(p ClientManagerParams) *ClientManager {
	config, err := p.CsConfigSvc.GetCsConfig(context.Background())
	if err != nil {
		return nil
	}

	return &ClientManager{
		lock:          &sync.RWMutex{}, // 需要加鎖防止等待用戶超出上限
		memberClients: &sync.Map{},     // roomId maps member client
		dispatcher:    p.Dispatcher,
		notifier:      p.Notifier,
		roomSvc:       p.RoomSvc,
		memberSvc:     p.MemberSvc,
		csConfig:      config,
		lua:           p.Lua,
		memberPool:    &sync.Pool{New: func() any { return &MemberClient{} }},
		staffPool:     &sync.Pool{New: func() any { return &StaffClient{} }},
	}
}

func checkPendingTimeout(member *MemberClient) {
	manager := member.Manager
	d := time.Duration(manager.csConfig.MemberPendingExpire) * time.Minute
	select {
	case <-time.After(d):
		manager.notifier.RoomClosed(member)
		time.Sleep(1 * time.Second)
		if err := manager.Unregister(member); err != nil {
			logger.Logger.Errorf("unregister staff client error: %s", err)
			return
		}
		// remove member token
		if err := manager.lua.RemoveMemberToken(context.Background(), fmt.Sprintf("%d-%s", member.SiteID, member.Name)); err != nil {
			logger.Logger.Errorf("timeout -> clear member token error: %s\n", err)
			return
		}
	case <-member.Sending:
		return
	}
}
