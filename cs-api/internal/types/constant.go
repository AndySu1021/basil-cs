package types

type ConstantType int8

const (
	ConstantTypeFastMessage ConstantType = iota + 1
	ConstantTypeCsConfig
)

type ConstantKey string

const (
	ConstantKeyFastMessage = "FastMessage"
	ConstantKeyCsConfig    = "CsConfig"
)

type CsConfig struct {
	MaxMember           int64  `json:"max_member"`
	MemberPendingExpire int64  `json:"member_pending_expire"`
	GreetingText        string `json:"greeting_text"`
}
