package mock

import (
	iface "cs-api/internal/interface"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewLuaScript(t *testing.T) iface.ILuaScript {
	m := gomock.NewController(t)
	mock := NewMockILuaScript(m)

	mock.EXPECT().SetToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().RemoveMemberToken(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().RemoveStaffToken(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().StaffAssignRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().StaffRegister(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().StaffCloseRoom(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().StaffTransferRoom(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return mock
}
