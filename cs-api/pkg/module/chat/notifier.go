package chat

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"encoding/json"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/AndySu1021/go-util/logger"
	"time"
)

type Notifier struct {
	redis  ifaceTool.IRedis
	msgSvc iface.IMessageService
}

func NewNotifier(redis ifaceTool.IRedis, msgSvc iface.IMessageService) *Notifier {
	return &Notifier{
		redis:  redis,
		msgSvc: msgSvc,
	}
}

func (n *Notifier) MemberJoin(staffId int64, member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeMemberJoin,
		Content:   member.Name,
		Extra:     types.JSONField{"member_id": member.ID},
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			StaffID: &staffId,
			RoomID:  &member.RoomID,
			Message: &message,
		},
	}

	n.send(info, false)
}

// StaffTyping 通知用戶客服正在輸入
func (n *Notifier) StaffTyping(roomId int64, staffName string) {
	message := model.Message{
		RoomID:    roomId,
		OpType:    types.MessageOpTypeStaffTyping,
		Content:   staffName,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			RoomID:  &roomId,
			Message: &message,
		},
	}

	n.send(info, false)
}

// StaffBroadcast 廣播客服消息
func (n *Notifier) StaffBroadcast(clientMessage ClientMessage, staff *StaffClient) {
	message := model.Message{
		RoomID:      clientMessage.RoomID,
		OpType:      types.MessageOpTypeChatMessage,
		SenderType:  types.MessageSenderTypeStaff,
		SenderID:    staff.ID,
		SenderName:  staff.Name,
		ContentType: clientMessage.ContentType,
		Content:     clientMessage.Content,
		Ts:          time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			RoomID:  &clientMessage.RoomID,
			Message: &message,
		},
	}

	byteMessage, _ := json.Marshal(message)
	staff.SendChan <- byteMessage

	n.send(info, true)
}

func (n *Notifier) MemberBroadcast(clientMessage ClientMessage, member *MemberClient) {
	message := model.Message{
		RoomID:      clientMessage.RoomID,
		OpType:      types.MessageOpTypeChatMessage,
		SenderType:  types.MessageSenderTypeMember,
		SenderID:    member.ID,
		SenderName:  member.Name,
		ContentType: clientMessage.ContentType,
		Content:     clientMessage.Content,
		Ts:          time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			StaffID: &member.StaffID,
			RoomID:  &clientMessage.RoomID,
			Message: &message,
		},
	}

	byteMessage, _ := json.Marshal(message)
	member.SendChan <- byteMessage

	n.send(info, true)
}

// NoStaff 通知客戶當前沒有客服可以服務
func (n *Notifier) NoStaff(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeNoStaff,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	byteMessage, _ := json.Marshal(message)
	member.SendChan <- byteMessage
}

// RoomClosed 通知客戶房間已關閉
func (n *Notifier) RoomClosed(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeRoomClosed,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	byteMessage, _ := json.Marshal(message)
	member.SendChan <- byteMessage
}

// SendScore 通知用戶評分
func (n *Notifier) SendScore(roomId int64, staff *StaffClient) {
	message := model.Message{
		RoomID:    roomId,
		OpType:    types.MessageOpTypeSendScore,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			RoomID:  &roomId,
			Message: &message,
		},
	}

	byteMessage, _ := json.Marshal(message)
	staff.SendChan <- byteMessage

	n.send(info, true)
}

// MemberScored 通知客服用戶已完成評分
func (n *Notifier) MemberScored(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeCompleteScore,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			StaffID: &member.StaffID,
			RoomID:  &member.RoomID,
			Message: &message,
		},
	}

	n.send(info, true)
}

// RoomAccepted 通知客戶客服已接受並進入房間
func (n *Notifier) RoomAccepted(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeRoomAccepted,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	byteMessage, _ := json.Marshal(message)
	member.SendChan <- byteMessage
}

// Greeting 通知客戶問候語
func (n *Notifier) Greeting(content string, member *MemberClient) {
	if content == "" {
		return
	}

	message := model.Message{
		RoomID:      member.RoomID,
		OpType:      types.MessageOpTypeGreeting,
		SenderType:  types.MessageSenderTypeStaff,
		SenderID:    0,
		SenderName:  "Basil 客服",
		ContentType: types.MessageContentTypeText,
		Content:     content,
		Ts:          time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	byteMessage, _ := json.Marshal(message)
	member.SendChan <- byteMessage
}

// RoomTransferred 通知用戶已被轉接
func (n *Notifier) RoomTransferred(roomId int64) {
	message := model.Message{
		RoomID:    roomId,
		OpType:    types.MessageOpTypeRoomTransferred,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	info := pkg.ClientEventInfo{
		Event: pkg.ClientEventMessageReceived,
		Payload: pkg.ClientEventPayload{
			RoomID:  &roomId,
			Message: &message,
		},
	}

	n.send(info, false)
}

func (n *Notifier) send(info pkg.ClientEventInfo, needStore bool) {
	if needStore {
		params := model.CreateMessageParams{
			RoomID:      info.Payload.Message.RoomID,
			OpType:      info.Payload.Message.OpType,
			SenderType:  info.Payload.Message.SenderType,
			SenderID:    info.Payload.Message.SenderID,
			SenderName:  info.Payload.Message.SenderName,
			ContentType: info.Payload.Message.ContentType,
			Content:     info.Payload.Message.Content,
			Extra:       info.Payload.Message.Extra,
			Ts:          info.Payload.Message.Ts,
			CreatedAt:   info.Payload.Message.CreatedAt,
		}
		if params.Extra == nil {
			params.Extra = types.JSONField{}
		}
		if err := n.msgSvc.CreateMessage(context.Background(), params); err != nil {
			logger.Logger.Errorf("create message error: %s\n", err)
		}
	}

	byteInfo, _ := json.Marshal(info)
	err := n.redis.Publish(context.Background(), types.RedisKeyEventClient, byteInfo)
	if err != nil {
		logger.Logger.Errorf("publish message to redis error: %s", err.Error())
		return
	}
}
