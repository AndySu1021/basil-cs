package message

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/internal/interface"
	"cs-api/internal/types"
	"reflect"
	"testing"
)

func Test_service_CreateMessage(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateMessageParams
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
				params: model.CreateMessageParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateMessage(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListRoomMessage(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMessages []model.Message
		wantErr      bool
	}{
		{
			name:   "normal test - member",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.ListMemberRoomMessageParams{},
			},
			wantMessages: make([]model.Message, 0),
			wantErr:      false,
		},
		{
			name:   "normal test - staff",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.ListStaffRoomMessageParams{},
			},
			wantMessages: make([]model.Message, 0),
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotMessages, err := s.ListRoomMessage(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRoomMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMessages, tt.wantMessages) {
				t.Errorf("ListRoomMessage() gotMessages = %v, want %v", gotMessages, tt.wantMessages)
			}
		})
	}
}

func Test_service_ListMessage(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params types.ListMessageParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []types.ListMessageRow
		want1   int64
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: types.ListMessageParams{},
			},
			want:    make([]types.ListMessageRow, 0),
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, got1, err := s.ListMessage(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
