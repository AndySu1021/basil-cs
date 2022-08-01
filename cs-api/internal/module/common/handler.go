package common

import (
	iface "cs-api/internal/interface"
	"fmt"
	ginTool "github.com/AndySu1021/go-util/gin"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
)

type handler struct {
	commonSvc iface.ICommonService
	redis     ifaceTool.IRedis
	storage   ifaceTool.IStorage
}

func (h *handler) GetPanelData(c *gin.Context) {
	data, err := h.commonSvc.GetPanelData(c.Request.Context())
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, data)
}

type UploadFileParams struct {
	Type string `form:"type" binding:"required,oneof=member staff"`
}

func (h *handler) UploadFile(c *gin.Context) {
	token := c.GetHeader("X-Token")

	var req UploadFileParams
	if err := c.ShouldBind(&req); err != nil {
		ginTool.Error(c, err)
		return
	}

	redisKey := fmt.Sprintf("token:%s:%s", req.Type, token)
	_, err := h.redis.Get(c.Request.Context(), redisKey).Result()
	if err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tmp, err := file.Open()
	if err != nil {
		ginTool.Error(c, err)
		return
	}
	defer tmp.Close()

	url, err := h.storage.Upload(c.Request.Context(), tmp, "upload", file.Filename)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, url)
}

func NewHandler(commonSvc iface.ICommonService, redis ifaceTool.IRedis, storage ifaceTool.IStorage) *handler {
	return &handler{
		commonSvc: commonSvc,
		redis:     redis,
		storage:   storage,
	}
}
