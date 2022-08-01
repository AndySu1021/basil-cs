package module

import (
	"cs-api/internal/module/auth"
	"cs-api/internal/module/chat"
	"cs-api/internal/module/common"
	"cs-api/internal/module/cs_config"
	"cs-api/internal/module/faq"
	"cs-api/internal/module/fast_reply"
	"cs-api/internal/module/member"
	"cs-api/internal/module/message"
	"cs-api/internal/module/notice"
	"cs-api/internal/module/remind"
	"cs-api/internal/module/report"
	"cs-api/internal/module/role"
	"cs-api/internal/module/room"
	"cs-api/internal/module/site"
	"cs-api/internal/module/staff"
	"cs-api/internal/module/tag"
	"cs-api/internal/module/tool"
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		tool.NewRequestInstrument,
	),
	auth.Options,
	tag.Options,
	role.Options,
	notice.Options,
	remind.Options,
	member.Options,
	cs_config.Options,
	report.Options,
	staff.Options,
	message.Options,
	fast_reply.Options,
	room.Options,
	common.Options,
	chat.Options,
	faq.Options,
	site.Options,
)
