package tag

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"reflect"
	"testing"
)

func Test_service_ListTag(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListTagParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTags  []types.ListTagRow
		wantTotal int64
		wantErr   bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: types.ListTagParams{},
			},
			wantTags:  make([]types.ListTagRow, 0),
			wantTotal: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotTags, gotTotal, err := s.ListTag(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTags, tt.wantTags) {
				t.Errorf("ListTag() gotTags = %v, want %v", gotTags, tt.wantTags)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListTag() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_service_GetTag(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx   context.Context
		tagId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantTag model.GetTagRow
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:   context.Background(),
				tagId: 1,
			},
			wantTag: model.GetTagRow{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotTag, err := s.GetTag(tt.args.ctx, tt.args.tagId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTag, tt.wantTag) {
				t.Errorf("GetTag() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
		})
	}
}

func Test_service_CreateTag(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateTagParams
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
				params: model.CreateTagParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateTag(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateTag(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateTagParams
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
				params: model.UpdateTagParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateTag(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteTag(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx   context.Context
		tagId int64
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
				ctx:   context.Background(),
				tagId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteTag(tt.args.ctx, tt.args.tagId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListAvailableTag(t *testing.T) {
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
		want    []model.ListAvailableTagRow
		wantErr bool
	}{
		{
			name:    "normal test",
			fields:  fields{repo: mock.NewRepository(t)},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListAvailableTagRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListAvailableTag(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailableTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}
