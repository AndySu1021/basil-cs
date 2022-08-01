package notice

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"reflect"
	"testing"
)

func Test_service_ListNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListNoticeParams
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []types.ListNoticeRow
		wantTotal  int64
		wantErr    bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: types.ListNoticeParams{},
			},
			wantResult: make([]types.ListNoticeRow, 0),
			wantTotal:  0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotResult, gotTotal, err := s.ListNotice(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ListNotice() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListNotice() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_service_GetNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		noticeId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Notice
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				noticeId: 1,
			},
			want:    model.Notice{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetNotice(tt.args.ctx, tt.args.noticeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateNoticeParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.CreateNoticeParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateNotice(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateNoticeParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateNoticeParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateNotice(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		noticeId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				noticeId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteNotice(tt.args.ctx, tt.args.noticeId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetLatestNotice(t *testing.T) {
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
		want    model.GetLatestNoticeRow
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx: context.Background(),
			},
			want:    model.GetLatestNoticeRow{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetLatestNotice(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvailableNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableNotice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
