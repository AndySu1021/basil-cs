package chat

import (
	"cs-api/config"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"encoding/json"
	"fmt"
	ginTool "github.com/AndySu1021/go-util/gin"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/AndySu1021/go-util/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type handler struct {
	config  *config.Config
	redis   ifaceTool.IRedis
	manager *ClientManager
	roomSvc iface.IRoomService
}

type ChatParams struct {
	Type pkg.ClientType `form:"type" binding:"required"`
	SID  string         `form:"sid" binding:"required"`
}

func (h *handler) Chat(c *gin.Context) {
	var (
		req ChatParams
		ctx = c.Request.Context()
	)

	if err := c.BindQuery(&req); err != nil {
		ginTool.Error(c, err)
		return
	}

	// check token
	redisKey := fmt.Sprintf("token:%s:%s", req.Type, req.SID)
	payload, err := h.redis.Get(ctx, redisKey).Result()
	if err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	var clientInfo pkg.ClientInfo
	if err = json.Unmarshal([]byte(payload), &clientInfo); err != nil {
		ginTool.Error(c, err)
		return
	}

	// upgrade http protocol to websocket protocol and verify origin
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		for i := 0; i < len(h.config.WsOriginWhiteList); i++ {
			if h.config.WsOriginWhiteList[i] == origin {
				return true
			}
		}
		return false
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Logger.Errorf("upgrade connection error: %s", err)
		return
	}

	// dispatch member to original staff if room is not closed
	if clientInfo.Type == pkg.ClientTypeMember {
		room, err := h.roomSvc.GetMemberAvailableRoom(ctx, clientInfo.ID)
		if err != nil {
			logger.Logger.Errorf("get member available room error: %s", err)
			return
		}
		clientInfo.RoomID = room.ID
		clientInfo.StaffID = room.StaffID
	}

	clientInfo.Conn = conn

	h.manager.Register(clientInfo)
}

func NewHandler(config *config.Config, redis ifaceTool.IRedis, manager *ClientManager, roomSvc iface.IRoomService) *handler {
	return &handler{
		config:  config,
		redis:   redis,
		manager: manager,
		roomSvc: roomSvc,
	}
}
