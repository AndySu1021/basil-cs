package module

import (
	"cs-api/pkg/module/auth"
	"cs-api/pkg/module/chat"
	"cs-api/pkg/module/common"
	"cs-api/pkg/module/cs_config"
	"cs-api/pkg/module/faq"
	"cs-api/pkg/module/fast_reply"
	"cs-api/pkg/module/member"
	"cs-api/pkg/module/message"
	"cs-api/pkg/module/notice"
	"cs-api/pkg/module/remind"
	"cs-api/pkg/module/report"
	"cs-api/pkg/module/role"
	"cs-api/pkg/module/room"
	"cs-api/pkg/module/site"
	"cs-api/pkg/module/staff"
	"cs-api/pkg/module/tag"
	"cs-api/pkg/module/tool"
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
