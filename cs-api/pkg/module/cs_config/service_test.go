package cs_config

import (
	"context"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_GetCsConfig(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantConfig types.CsConfig
		wantErr    bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis: mockTool.NewRedis(t),
				repo:  mock.NewRepository(t),
			},
			args: args{ctx: context.Background()},
			wantConfig: types.CsConfig{
				MaxMember:           5,
				MemberPendingExpire: 0,
				GreetingText:        "",
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
			gotConfig, err := s.GetCsConfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCsConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("GetCsConfig() gotConfig = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

func Test_service_UpdateCsConfig(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		config  types.CsConfig
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
				repo:  mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				config:  types.CsConfig{},
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
			if err := s.UpdateCsConfig(tt.args.ctx, tt.args.staffId, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCsConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
