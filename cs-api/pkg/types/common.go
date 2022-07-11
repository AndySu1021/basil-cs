package types

import (
	"database/sql/driver"
	"encoding/json"
)

// Status 開啟狀態
type Status int8

const (
	StatusAll Status = iota
	StatusEnabled
	StatusDisabled
)

type Pagination struct {
	Page     int32 `json:"page" form:"page" binding:"required,gte=1"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required,gte=1"`
	Total    int64 `json:"total" form:"total"`
}

const (
	RequestCount       = "cs_request_count"
	RequestCountHelp   = "Total request count"
	RequestLatency     = "cs_request_latency"
	RequestLatencyHelp = "Total duration of request in microseconds"
)

type JSONField map[string]interface{}

func (f *JSONField) UnmarshalJSON(b []byte) error {
	d := make(map[string]interface{}, 0)
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	*f = d
	return nil
}

func (f *JSONField) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), f)
}

func (f JSONField) Value() (driver.Value, error) {
	return json.Marshal(&f)
}

type PanelData struct {
	StaffCount   int64 `json:"staff_count"`
	ServingCount int64 `json:"serving_count"`
	QueuingCount int64 `json:"queuing_count"`
}

const (
	RedisKeyStaffDispatch = "zset:staff:dispatch"
	RedisKeyEventClient   = "chan:event:client"
)
