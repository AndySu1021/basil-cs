package chat

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"encoding/json"
	"github.com/AndySu1021/go-util/logger"
	"github.com/gorilla/websocket"
)

type StaffClient struct {
	ID        int64
	Type      pkg.ClientType
	Name      string
	Socket    *websocket.Conn
	SendChan  chan []byte // 欲傳送的data
	RoomCount int64       // 當前服務尚未關閉的房間數量
	Status    ClientStatus
	Manager   *ClientManager
	Notifier  *Notifier
}

func (sc *StaffClient) GetID() int64 {
	return sc.ID
}

func (sc *StaffClient) GetName() string {
	return sc.Name
}

func (sc *StaffClient) GetType() pkg.ClientType {
	return sc.Type
}

func (sc *StaffClient) GetSenderType() types.MessageSenderType {
	return types.MessageSenderTypeStaff
}

func (sc *StaffClient) GetStatus() ClientStatus {
	return sc.Status
}

func (sc *StaffClient) GetSendChan() chan []byte {
	return sc.SendChan
}

func (sc *StaffClient) SocketRead() {
	defer func() {
		if err := sc.Manager.Unregister(sc); err != nil {
			logger.Logger.Errorf("unregister staff client error: %s", err)
			return
		}
	}()

	for {
		mt, message, err := sc.Socket.ReadMessage()
		if err != nil {
			logger.Logger.Errorf("staff client ws error: %s\n", err)
			break
		}

		// heart beat
		if mt == websocket.PingMessage {
			if err = sc.Socket.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
				logger.Logger.Errorf("staff client write pong message error: %s\n", err)
				return
			}
			continue
		}

		var tmp ClientMessage
		if err = json.Unmarshal(message, &tmp); err != nil {
			logger.Logger.Errorf("error: %s\n", err)
			break
		}

		switch tmp.OpType {
		case types.MessageOpTypeChatMessage:
			sc.Notifier.StaffBroadcast(tmp, sc)
		case types.MessageOpTypeStaffTyping:
			sc.Notifier.StaffTyping(tmp.RoomID, sc.Name)
		case types.MessageOpTypeSendScore:
			sc.Notifier.SendScore(tmp.RoomID, sc)
		}
	}
}

func (sc *StaffClient) SocketWrite() {
	defer func() {
		if err := sc.Manager.Unregister(sc); err != nil {
			logger.Logger.Errorf("unregister staff client error: %s", err)
			return
		}
	}()

	for {
		message, ok := <-sc.SendChan
		if !ok {
			if err := sc.Socket.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				logger.Logger.Errorf("staff client write close message error: %s\n", err)
			}
			return
		}

		if err := sc.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
			logger.Logger.Errorf("staff client write message error: %s\n", err)
			return
		}
	}
}

// Reset 初始化從 sync.Pool 中拿出的 StaffClient
func (sc *StaffClient) Reset(clientInfo pkg.ClientInfo, manager *ClientManager) error {
	roomCount, err := manager.roomSvc.GetStaffServingRoomCount(context.Background(), clientInfo.ID)
	if err != nil {
		logger.Logger.Errorf("get staff room count error: %s", err.Error())
		return err
	}

	sc.ID = clientInfo.ID
	sc.Type = pkg.ClientTypeStaff
	sc.Name = clientInfo.Name
	sc.Socket = clientInfo.Conn
	sc.SendChan = make(chan []byte)
	sc.RoomCount = roomCount
	sc.Status = ClientStatusOpen
	sc.Manager = manager
	sc.Notifier = manager.notifier

	return nil
}
