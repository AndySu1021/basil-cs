// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"cs-api/pkg/types"
	"github.com/shopspring/decimal"
)

// 常數資料表
type Constant struct {
	ID int64 `db:"id" json:"id"`
	// 常數類型 1快捷回覆 2客服配置
	Type types.ConstantType `db:"type" json:"type"`
	// 鍵
	Key types.ConstantKey `db:"key" json:"key"`
	// 值
	Value string `db:"value" json:"value"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 常見問題資料表
type Faq struct {
	ID int64 `db:"id" json:"id"`
	// 問題
	Question string `db:"question" json:"question"`
	// 答案
	Answer string `db:"answer" json:"answer"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 快捷回覆資料表
type FastReply struct {
	ID int64 `db:"id" json:"id"`
	// 分類ID(constantID)
	CategoryID int64 `db:"category_id" json:"category_id"`
	// 訊息標題
	Title string `db:"title" json:"title"`
	// 訊息內容
	Content string `db:"content" json:"content"`
	// 消息狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 會員資料表
type Member struct {
	ID int64 `db:"id" json:"id"`
	// 用戶類型 1一般用戶 2訪客
	Type types.MemberType `db:"type" json:"type"`
	// 站點ID
	SiteID int64 `db:"site_id" json:"site_id"`
	// 會員名稱
	Name string `db:"name" json:"name"`
	// 設備號
	DeviceID string `db:"device_id" json:"device_id"`
	// 真實姓名
	RealName string `db:"real_name" json:"real_name"`
	// 手機號
	Mobile string `db:"mobile" json:"mobile"`
	// 電子信箱
	Email string `db:"email" json:"email"`
	// 備註
	Note string `db:"note" json:"note"`
	// 會員狀態 1在線 2離線
	OnlineStatus types.MemberOnlineStatus `db:"online_status" json:"online_status"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 歷史訊息紀錄表
type Message struct {
	ID int64 `db:"id" json:"id"`
	// 房間ID
	RoomID int64 `db:"room_id" json:"room_id"`
	// 操作類型 0其他 1客服輸入中 2發送訊息
	OpType types.MessageOpType `db:"op_type" json:"op_type"`
	// 發送人類型 0系統 1會員 2客服
	SenderType types.MessageSenderType `db:"sender_type" json:"sender_type"`
	// 發送人ID
	SenderID int64 `db:"sender_id" json:"sender_id"`
	// 發送人名稱
	SenderName string `db:"sender_name" json:"sender_name"`
	// 消息類型 0其他 1文字 2圖片
	ContentType types.MessageContentType `db:"content_type" json:"content_type"`
	// 訊息內容
	Content string `db:"content" json:"content"`
	// 額外資訊
	Extra types.JSONField `db:"extra" json:"extra"`
	// 創建時間戳
	Ts int64 `db:"ts" json:"ts"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// 系統公告資料表
type Notice struct {
	ID int64 `db:"id" json:"id"`
	// 公告標題
	Title string `db:"title" json:"title"`
	// 公告內容
	Content string `db:"content" json:"content"`
	// 開始時間
	StartAt time.Time `db:"start_at" json:"start_at"`
	// 結束時間
	EndAt time.Time `db:"end_at" json:"end_at"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 後台提醒事項資料表
type Remind struct {
	ID int64 `db:"id" json:"id"`
	// 標題
	Title string `db:"title" json:"title"`
	// 內容
	Content string `db:"content" json:"content"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 每日客服報表
type ReportDailyStaff struct {
	ID int64 `db:"id" json:"id"`
	// 報表日期
	Date time.Time `db:"date" json:"date"`
	// 職員ID
	StaffID int64 `db:"staff_id" json:"staff_id"`
	// 總接待人數
	TotalMember int32 `db:"total_member" json:"total_member"`
	// 總評分人數
	TotalScoredMember int32 `db:"total_scored_member" json:"total_scored_member"`
	// 平均分數
	AverageScore decimal.Decimal `db:"average_score" json:"average_score"`
}

// 每日標籤報表
type ReportDailyTag struct {
	ID int64 `db:"id" json:"id"`
	// 報表日期
	Date time.Time `db:"date" json:"date"`
	// 標籤ID
	TagID int64 `db:"tag_id" json:"tag_id"`
	// 人數
	Count int32 `db:"count" json:"count"`
}

// 權限角色資料表
type Role struct {
	ID int64 `db:"id" json:"id"`
	// 角色名稱
	Name string `db:"name" json:"name"`
	// 角色權限
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 諮詢房紀錄表
type Room struct {
	ID int64 `db:"id" json:"id"`
	// 職員ID
	StaffID int64 `db:"staff_id" json:"staff_id"`
	// 會員ID
	MemberID int64 `db:"member_id" json:"member_id"`
	// 訪問IP
	Ip string `db:"ip" json:"ip"`
	// 來源 1-web 2-app
	Source types.RoomSource `db:"source" json:"source"`
	// 請求使用者代理
	UserAgent string `db:"user_agent" json:"user_agent"`
	// 標籤ID
	TagID int64 `db:"tag_id" json:"tag_id"`
	// 評分 1-5
	Score int32 `db:"score" json:"score"`
	// 客服房狀態 1等待中 2排隊中 3服務中 4已關閉
	Status types.RoomStatus `db:"status" json:"status"`
	// 開始諮詢時間
	AcceptedAt sql.NullTime `db:"accepted_at" json:"accepted_at"`
	// 關閉時間
	ClosedAt sql.NullTime `db:"closed_at" json:"closed_at"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 站點資料表
type Site struct {
	ID int64 `db:"id" json:"id"`
	// 名稱
	Name string `db:"name" json:"name"`
	// 代號
	Code string `db:"code" json:"code"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 職員資料表
type Staff struct {
	ID int64 `db:"id" json:"id"`
	// 角色ID
	RoleID int64 `db:"role_id" json:"role_id"`
	// 職員姓名
	Name string `db:"name" json:"name"`
	// 用戶名
	Username string `db:"username" json:"username"`
	// 密碼
	Password string `db:"password" json:"password"`
	// 大頭貼
	Avatar string `db:"avatar" json:"avatar"`
	// 職員狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 職員服務狀態 1關閉 2服務中 3閒置
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	// 職員線上狀態 1在線 2離線
	OnlineStatus types.StaffOnlineStatus `db:"online_status" json:"online_status"`
	// 上次登入時間
	LastLoginTime sql.NullTime `db:"last_login_time" json:"last_login_time"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 標籤資料表
type Tag struct {
	ID int64 `db:"id" json:"id"`
	// 標籤名稱
	Name string `db:"name" json:"name"`
	// 標籤狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}