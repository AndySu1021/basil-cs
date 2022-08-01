package role

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_ListRole(t *testing.T) {
	type fields struct {
		repo  iface.IRepository
		redis ifaceTool.IRedis
	}
	type args struct {
		ctx    context.Context
		params types.ListRoleParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []types.ListRoleRow
		want1   int64
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				params: types.ListRoleParams{},
			},
			want:    make([]types.ListRoleRow, 0),
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:  tt.fields.repo,
				redis: tt.fields.redis,
			}
			got, got1, err := s.ListRole(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRole() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListRole() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_GetRole(t *testing.T) {
	type fields struct {
		repo  iface.IRepository
		redis ifaceTool.IRedis
	}
	type args struct {
		ctx    context.Context
		roleId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Role
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				roleId: 1,
			},
			want:    model.Role{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:  tt.fields.repo,
				redis: tt.fields.redis,
			}
			got, err := s.GetRole(tt.args.ctx, tt.args.roleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateRole(t *testing.T) {
	type fields struct {
		repo  iface.IRepository
		redis ifaceTool.IRedis
	}
	type args struct {
		ctx    context.Context
		params model.CreateRoleParams
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				params: model.CreateRoleParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:  tt.fields.repo,
				redis: tt.fields.redis,
			}
			if err := s.CreateRole(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateRole(t *testing.T) {
	type fields struct {
		repo  iface.IRepository
		redis ifaceTool.IRedis
	}
	type args struct {
		ctx    context.Context
		params model.UpdateRoleParams
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateRoleParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:  tt.fields.repo,
				redis: tt.fields.redis,
			}
			if err := s.UpdateRole(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteRole(t *testing.T) {
	type fields struct {
		repo  iface.IRepository
		redis ifaceTool.IRedis
	}
	type args struct {
		ctx    context.Context
		roleId int64
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				roleId: 2,
			},
			wantErr: false,
		},
		{
			name:   "delete error",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				roleId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:  tt.fields.repo,
				redis: tt.fields.redis,
			}
			if err := s.DeleteRole(tt.args.ctx, tt.args.roleId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
