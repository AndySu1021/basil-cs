package common

import (
	"context"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"reflect"
	"testing"
)

func Test_service_GetPanelData(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData types.PanelData
		wantErr  bool
	}{
		{
			name:     "normal test",
			fields:   fields{repo: mock.NewRepository(t)},
			args:     args{ctx: context.Background()},
			wantData: types.PanelData{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotData, err := s.GetPanelData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPanelData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetPanelData() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
