package internal

import (
	"cs-api/db/model"
	"cs-api/internal/types"
	"github.com/gorilla/websocket"
)

type ClientType string

const (
	ClientTypeStaff  ClientType = "staff"
	ClientTypeMember ClientType = "member"
)

type ClientInfo struct {
	ID          int64           `json:"id,omitempty"`
	Type        ClientType      `json:"type,omitempty"`
	Name        string          `json:"name,omitempty"`
	Username    string          `json:"username,omitempty"`    // staff only
	RoleID      int64           `json:"role_id,omitempty"`     // staff only
	Permissions []string        `json:"permissions,omitempty"` // staff only
	SiteID      int64           `json:"site_id,omitempty"`     // member only
	RoomID      int64           `json:"room_id,omitempty"`     // member only
	StaffID     int64           `json:"staff_id,omitempty"`    // member only
	Token       string          `json:"token,omitempty"`
	Conn        *websocket.Conn `json:"conn,omitempty"`
}

type ClientEvent string

const (
	ClientEventCloseRoom       ClientEvent = "CloseRoom"
	ClientEventAcceptRoom      ClientEvent = "AcceptRoom"
	ClientEventUpdateConfig    ClientEvent = "UpdateConfig"
	ClientEventTransferRoom    ClientEvent = "TransferRoom"
	ClientEventMessageReceived ClientEvent = "MessageReceived"
)

type ClientEventPayload struct {
	StaffID  *int64          `json:"staff_id,omitempty"`
	RoomID   *int64          `json:"room_id,omitempty"`
	CsConfig *types.CsConfig `json:"cs_config,omitempty"`
	Message  *model.Message  `json:"message"`
}

type ClientEventInfo struct {
	Event   ClientEvent        `json:"event"`
	Payload ClientEventPayload `json:"payload"`
}

type FastReplyCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type FastReplyGroupItem struct {
	Category FastReplyCategory                   `json:"category"`
	Items    []model.GetAllAvailableFastReplyRow `json:"items"`
}

type DailyTagReportColumn struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
