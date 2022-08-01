package site

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"reflect"
	"testing"
)

func Test_service_ListSite(t *testing.T) {
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
		want    []model.ListSiteRow
		wantErr bool
	}{
		{
			name:    "normal test",
			fields:  fields{repo: mock.NewRepository(t)},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListSiteRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListSite(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSite() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetSite(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		siteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.GetSiteRow
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				siteId: 1,
			},
			want:    model.GetSiteRow{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetSite(tt.args.ctx, tt.args.siteId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSite() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateSite(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateSiteParams
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
				params: model.CreateSiteParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateSite(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateSite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateSite(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateSiteParams
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
				params: model.UpdateSiteParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateSite(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteSite(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		siteId int64
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
				siteId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteSite(tt.args.ctx, tt.args.siteId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetSiteByCode(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.GetSiteByCodeRow
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:  context.Background(),
				code: "code",
			},
			want:    model.GetSiteByCodeRow{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetSiteByCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSiteByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSiteByCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
