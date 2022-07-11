package mock

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewMemberService(t *testing.T) iface.IMemberService {
	m := gomock.NewController(t)
	mock := NewMockIMemberService(m)

	mock.EXPECT().GetOrCreateMember(gomock.Any(), "name", "deviceId", gomock.Any()).AnyTimes().Return(model.Member{ID: 1}, nil)
	mock.EXPECT().GetOrCreateMember(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(model.Member{ID: 2}, nil)

	return mock
}

func NewSiteService(t *testing.T) iface.ISiteService {
	m := gomock.NewController(t)
	mock := NewMockISiteService(m)

	mock.EXPECT().GetSiteByCode(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetSiteByCodeRow{ID: 1}, nil)

	return mock
}
