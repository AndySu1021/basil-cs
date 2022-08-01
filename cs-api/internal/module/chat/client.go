package chat

import (
	"cs-api/internal"
	"cs-api/internal/types"
)

type ClientStatus int8

const (
	ClientStatusOpen ClientStatus = iota + 1
	ClientStatusClosed
)

type ClientMessage struct {
	RoomID      int64                    `json:"room_id"`
	OpType      types.MessageOpType      `json:"op_type"`
	ContentType types.MessageContentType `json:"content_type"`
	Content     string                   `json:"content"`
}

type Client interface {
	GetID() int64
	GetName() string
	GetType() internal.ClientType
	GetSenderType() types.MessageSenderType
	GetStatus() ClientStatus
	GetSendChan() chan []byte
	SocketRead()
	SocketWrite()
	Reset(clientInfo internal.ClientInfo, manager *ClientManager) error
}

// TODO: Maybe it should add the heart beat to check if connection is closed
