package remind

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"reflect"
	"testing"
)

func Test_service_ListRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListRemindParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []types.ListRemindRow
		want1   int64
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: types.ListRemindParams{},
			},
			want:    make([]types.ListRemindRow, 0),
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, got1, err := s.ListRemind(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRemind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRemind() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListRemind() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_GetRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		remindId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Remind
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				remindId: 1,
			},
			want:    model.Remind{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetRemind(tt.args.ctx, tt.args.remindId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRemind() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateRemindParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.CreateRemindParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateRemind(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateRemind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateRemindParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateRemindParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateRemind(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRemind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		remindId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				remindId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteRemind(tt.args.ctx, tt.args.remindId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRemind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListActiveRemind(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.ListActiveRemindRow
		wantErr bool
	}{
		{
			name:    "normal test",
			fields:  fields{repo: mock.NewRepository(t)},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListActiveRemindRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListActiveRemind(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListActiveRemind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListActiveRemind() got = %v, want %v", got, tt.want)
			}
		})
	}
}
