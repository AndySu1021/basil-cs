package chat

import (
	"cs-api/pkg"
	"cs-api/pkg/types"
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
	GetType() pkg.ClientType
	GetSenderType() types.MessageSenderType
	GetStatus() ClientStatus
	GetSendChan() chan []byte
	SocketRead()
	SocketWrite()
	Reset(clientInfo pkg.ClientInfo, manager *ClientManager) error
}

// TODO: Maybe it should add the heart beat to check if connection is closed
