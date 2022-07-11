package room

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	ginTool "github.com/AndySu1021/go-util/gin"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_CreateRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx       context.Context
		source    types.RoomSource
		userAgent string
		memberId  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantRoomId int64
		wantErr    bool
	}{
		{
			name: "normal test - room exists",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:       context.WithValue(context.Background(), ginTool.ContextKeyClientIp, "127.0.0.1"),
				source:    1,
				userAgent: "agent",
				memberId:  1,
			},
			wantRoomId: 1,
			wantErr:    false,
		},
		{
			name: "normal test - room doesn't exist",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:       context.WithValue(context.Background(), ginTool.ContextKeyClientIp, "127.0.0.1"),
				source:    1,
				userAgent: "agent",
				memberId:  2,
			},
			wantRoomId: 1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			gotRoomId, err := s.CreateRoom(tt.args.ctx, tt.args.source, tt.args.userAgent, tt.args.memberId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRoomId != tt.wantRoomId {
				t.Errorf("CreateRoom() gotRoomId = %v, want %v", gotRoomId, tt.wantRoomId)
			}
		})
	}
}

func Test_service_AcceptRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		roomId  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test - pending room",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  1,
			},
			wantErr: false,
		},
		{
			name: "normal test - closed room",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			if err := s.AcceptRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId); (err != nil) != tt.wantErr {
				t.Errorf("AcceptRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_CloseRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		roomId  int64
		tagId   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test - serving room",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  2,
				tagId:   1,
			},
			wantErr: false,
		},
		{
			name: "normal test - pending room",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  1,
				tagId:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			if err := s.CloseRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId, tt.args.tagId); (err != nil) != tt.wantErr {
				t.Errorf("CloseRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateRoomScore(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		roomId int64
		score  int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				roomId: 1,
				score:  5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateRoomScore(tt.args.ctx, tt.args.roomId, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoomScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListRoomParams
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []types.ListRoomRow
		wantTotal  int64
		wantErr    bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				params: types.ListRoomParams{},
			},
			wantResult: make([]types.ListRoomRow, 0),
			wantTotal:  0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			gotResult, gotTotal, err := s.ListRoom(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ListRoom() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListRoom() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_service_ListStaffServingRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []types.StaffRoom
		wantErr    bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			wantResult: make([]types.StaffRoom, 0),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			gotResult, err := s.ListStaffServingRoom(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStaffRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ListStaffRoom() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_service_ListQueuingRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.ListQueuingRoomRow
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListQueuingRoomRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			got, err := s.ListQueuingRoom(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPendingRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListPendingRoom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetStaffServingRoomCount(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			got, err := s.GetStaffServingRoomCount(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffRooms() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_TransferRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx       context.Context
		staffId   int64
		roomId    int64
		toStaffId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    2,
				toStaffId: 2,
			},
			wantErr: false,
		},
		{
			name: "staff not available",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    1,
				toStaffId: 1,
			},
			wantErr: true,
		},
		{
			name: "no need to transfer",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    3,
				toStaffId: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			if err := s.TransferRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId, tt.args.toStaffId); (err != nil) != tt.wantErr {
				t.Errorf("TransferRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetMemberAvailableRoom(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx      context.Context
		memberId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Room
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:      context.Background(),
				memberId: 1,
			},
			want:    model.Room{ID: 1, MemberID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			got, err := s.GetMemberAvailableRoom(tt.args.ctx, tt.args.memberId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberAvailableRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemberAvailableRoom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetRoomInfo(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		roomId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.GetRoomInfoRow
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				roomId: 1,
			},
			want:    model.GetRoomInfoRow{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			got, err := s.GetRoomInfo(tt.args.ctx, tt.args.roomId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoomInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoomInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateRoomStatus(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		lua   iface.ILuaScript
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateRoomStatusParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				lua:   mock.NewLuaScript(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateRoomStatusParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				lua:   tt.fields.lua,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateRoomStatus(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoomStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
