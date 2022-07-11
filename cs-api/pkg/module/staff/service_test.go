package staff

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_ListStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListStaffParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []types.ListStaffRow
		want1   int64
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				params: types.ListStaffParams{},
			},
			want:    make([]types.ListStaffRow, 0),
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			got, got1, err := s.ListStaff(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListStaff() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListStaff() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_GetStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantStaff model.GetStaffRow
		wantErr   bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			wantStaff: model.GetStaffRow{OnlineStatus: 1, ServingStatus: 1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			gotStaff, err := s.GetStaff(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaff, tt.wantStaff) {
				t.Errorf("GetStaff() gotStaff = %v, want %v", gotStaff, tt.wantStaff)
			}
		})
	}
}

func Test_service_CreateStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateStaffParams
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
				params: model.CreateStaffParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.CreateStaff(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params interface{}
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
				params: model.UpdateStaffParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateStaff(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
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
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.DeleteStaff(tt.args.ctx, tt.args.staffId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateStaffServingStatus(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		status  types.StaffServingStatus
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
				ctx:     context.Background(),
				staffId: 1,
				status:  types.StaffServingStatusServing,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateStaffServingStatus(tt.args.ctx, tt.args.staffId, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaffServingStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListAvailableStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
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
		want    []model.ListAvailableStaffRow
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			want:    make([]model.ListAvailableStaffRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			got, err := s.ListAvailableStaff(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailableStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
