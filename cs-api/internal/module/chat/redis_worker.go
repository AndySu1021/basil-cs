package chat

import (
	"context"
	"cs-api/internal"
	"cs-api/internal/types"
	"encoding/json"
	"github.com/AndySu1021/go-util/helper"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/AndySu1021/go-util/logger"
	"reflect"
	"time"
)

func InitRedisSubscriber(redis iface.IRedis, worker *RedisWorker) {
	subscriber := redis.Subscribe(context.Background(), types.RedisKeyEventClient)

	var info internal.ClientEventInfo

	go func() {
		defer helper.Recover(context.Background())
		for {
			msg, err := subscriber.ReceiveMessage(context.Background())
			if err != nil {
				logger.Logger.Errorf("error: %s\n", err.Error())
				return
			}

			if err = json.Unmarshal([]byte(msg.Payload), &info); err != nil {
				logger.Logger.Errorf("error: %s\n", err.Error())
				return
			}

			worker.Handle(info)
		}
	}()
}

type RedisWorker struct {
	handler *EventHandler
	methods map[internal.ClientEvent]reflect.Method
}

func NewRedisWorker(manager *ClientManager, dispatcher *StaffDispatcher, notifier *Notifier) *RedisWorker {
	worker := &RedisWorker{
		handler: &EventHandler{manager: manager, dispatcher: dispatcher, notifier: notifier},
		methods: make(map[internal.ClientEvent]reflect.Method),
	}
	worker.registerMethods()
	return worker
}

func (r *RedisWorker) Handle(info internal.ClientEventInfo) {
	f := r.methods[info.Event].Func
	f.Call([]reflect.Value{reflect.ValueOf(r.handler), reflect.ValueOf(info)})
}

func (r *RedisWorker) registerMethods() {
	t := reflect.TypeOf(r.handler)
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		r.methods[internal.ClientEvent(name)] = t.Method(i)
	}
}

type EventHandler struct {
	manager    *ClientManager
	dispatcher *StaffDispatcher
	notifier   *Notifier
}

func (r *EventHandler) CloseRoom(eventInfo internal.ClientEventInfo) {
	roomId := *eventInfo.Payload.RoomID
	if member := r.manager.GetMember(roomId); member != nil {
		r.notifier.RoomClosed(member)
		time.Sleep(100 * time.Millisecond)
		r.manager.CloseRoom(roomId)
	}
}

func (r *EventHandler) AcceptRoom(eventInfo internal.ClientEventInfo) {
	roomId := *eventInfo.Payload.RoomID
	if member := r.manager.GetMember(roomId); member == nil {
		return
	}
	member := r.manager.GetMember(roomId)
	member.StaffID = *eventInfo.Payload.StaffID
	if err := r.dispatcher.assignRoom(*eventInfo.Payload.StaffID); err != nil {
		logger.Logger.Errorf("redis worker accept room error: %s", err.Error())
		return
	}
	r.notifier.RoomAccepted(member)
}

func (r *EventHandler) UpdateConfig(eventInfo internal.ClientEventInfo) {
	config := *eventInfo.Payload.CsConfig
	r.manager.csConfig = config
	r.dispatcher.setMaxMember(config.MaxMember)
}

func (r *EventHandler) TransferRoom(eventInfo internal.ClientEventInfo) {
	roomId := *eventInfo.Payload.RoomID
	staffId := *eventInfo.Payload.StaffID
	if member := r.manager.GetMember(roomId); member != nil {
		member.StaffID = staffId
		r.notifier.RoomTransferred(roomId)
		r.notifier.MemberJoin(staffId, member)
	}
}

func (r *EventHandler) MessageReceived(eventInfo internal.ClientEventInfo) {
	roomId := *eventInfo.Payload.RoomID
	switch eventInfo.Payload.Message.OpType {
	case types.MessageOpTypeChatMessage:
		if eventInfo.Payload.Message.SenderType == types.MessageSenderTypeStaff {
			if member := r.manager.GetMember(roomId); member != nil {
				byteMessage, _ := json.Marshal(*eventInfo.Payload.Message)
				member.SendChan <- byteMessage
			}
		} else if eventInfo.Payload.Message.SenderType == types.MessageSenderTypeMember {
			staffId := *eventInfo.Payload.StaffID
			if staff := r.dispatcher.getStaff(staffId); staff != nil {
				byteMessage, _ := json.Marshal(*eventInfo.Payload.Message)
				staff.SendChan <- byteMessage
			}
		}
	case types.MessageOpTypeStaffTyping,
		types.MessageOpTypeSendScore,
		types.MessageOpTypeRoomAccepted,
		types.MessageOpTypeRoomTransferred:
		if member := r.manager.GetMember(roomId); member != nil {
			byteMessage, _ := json.Marshal(*eventInfo.Payload.Message)
			member.SendChan <- byteMessage
		}
	case types.MessageOpTypeMemberJoin,
		types.MessageOpTypeCompleteScore:
		staffId := *eventInfo.Payload.StaffID
		if staff := r.dispatcher.getStaff(staffId); staff != nil {
			byteMessage, _ := json.Marshal(*eventInfo.Payload.Message)
			staff.SendChan <- byteMessage
		}
	}
}
