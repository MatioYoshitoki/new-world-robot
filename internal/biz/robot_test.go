package biz

import (
	"context"
	"math/rand"
	"net/http"
	v1 "new-world-robot/api/new-world-robot/v1"
	"new-world-robot/internal/conf"
	"new-world-robot/pkg/tutils"
	"testing"
	"time"
)

func TestRobot_auth(t *testing.T) {
	bs := tutils.LoadBs()
	type fields struct {
		Id                                 int64
		ticker                             time.Ticker
		hc                                 *http.Client
		rd                                 *rand.Rand
		memory                             *v1.RobotMemory
		bs                                 *conf.Bootstrap
		defaultRateMap                     []func(ctx context.Context) error
		canCreateFishRateMap               []func(ctx context.Context) error
		canExpandParkingRateMap            []func(ctx context.Context) error
		bothCreateFishExpandParkingRateMap []func(ctx context.Context) error
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				bs: bs,
				hc: &http.Client{},
				rd: rand.New(rand.NewSource(time.Now().UnixNano())),
				memory: &v1.RobotMemory{
					Auth: &v1.RobotMemory_Auth{
						Face: "insulate",
						FKey: "e10adc3949ba59abbe56e057f20f883e",
						Uid:  "1001",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				hc:     tt.fields.hc,
				rd:     tt.fields.rd,
				memory: tt.fields.memory,
				bs:     tt.fields.bs,
			}
			got, err := r.auth()
			if (err != nil) != tt.wantErr {
				t.Errorf("auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Error(got.AccessToken)
		})
	}
}
