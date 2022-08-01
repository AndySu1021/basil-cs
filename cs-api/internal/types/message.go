package types

type MessageOpType int8

const (
	// MessageOpTypeGreeting 招呼語
	MessageOpTypeGreeting MessageOpType = iota
	// MessageOpTypeStaffTyping 客服輸入中
	MessageOpTypeStaffTyping
	// MessageOpTypeChatMessage 收到聊天訊息
	MessageOpTypeChatMessage
	// MessageOpTypeSendScore 發送評分請求
	MessageOpTypeSendScore
	// MessageOpTypeCompleteScore 用戶完成評分
	MessageOpTypeCompleteScore
	// MessageOpTypeMemberJoin 會員加入房間
	MessageOpTypeMemberJoin
	// MessageOpTypeNoStaff 無客服可以服務客戶
	MessageOpTypeNoStaff
	// MessageOpTypeRoomClosed 關閉諮詢房
	MessageOpTypeRoomClosed
	// MessageOpTypeRoomAccepted 客服開始諮詢等待中諮詢房
	MessageOpTypeRoomAccepted
	// MessageOpTypeRoomTransferred 轉接諮詢房
	MessageOpTypeRoomTransferred
)

type MessageSenderType int8

const (
	MessageSenderTypeSystem MessageSenderType = iota
	MessageSenderTypeMember
	MessageSenderTypeStaff
)

type MessageContentType int8

const (
	MessageContentTypeElse MessageContentType = iota
	MessageContentTypeText
	MessageContentTypeImage
)

type ListMessageParams struct {
	RoomID  int64  `form:"room_id" binding:""`
	StaffID int64  `form:"staff_id" binding:""`
	Content string `form:"content" binding:""`
	Pagination
}

type ListMessageRow struct {
	ID          int64              `db:"id" json:"id"`
	SenderType  MessageSenderType  `db:"sender_type" json:"sender_type"`
	SenderName  string             `db:"sender_name" json:"sender_name"`
	ContentType MessageContentType `db:"content_type" json:"content_type"`
	Content     string             `db:"content" json:"content"`
	Ts          int64              `db:"ts" json:"ts"`
}
