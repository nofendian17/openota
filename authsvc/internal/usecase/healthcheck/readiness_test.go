package healthcheck

import (
	"context"
	"errors"
	"fmt"
	"github.com/nofendian17/openota/authsvc/internal/entity"
	"testing"
	"time"

	"github.com/nofendian17/openota/authsvc/internal/config"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
	mockCacheClient "github.com/nofendian17/openota/authsvc/internal/mocks/infra/cache"
	mockDB "github.com/nofendian17/openota/authsvc/internal/mocks/infra/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_useCase_Readiness(t *testing.T) {
	cfg := config.New()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		cacheErr error
		dbErr    error
		want     *entity.ReadinessResponse
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			cacheErr: nil,
			want: &entity.ReadinessResponse{
				Status: statusUP,
				Checks: []entity.Check{
					{
						Name:   "mysql",
						Status: statusUP,
					},
					{
						Name:   "Redis",
						Status: statusUP,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "failed - ping to redis got error",
			args: args{
				ctx: context.Background(),
			},
			cacheErr: errors.New("error"),
			want: &entity.ReadinessResponse{
				Status: statusDown,
				Checks: []entity.Check{
					{
						Name:   "mysql",
						Status: statusUP,
					},
					{
						Name:   "Redis",
						Status: statusDown,
						Error:  errors.New("error").Error(),
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "failed - ping to database got error",
			args: args{
				ctx: context.Background(),
			},
			dbErr:    errors.New("error"),
			cacheErr: nil,
			want: &entity.ReadinessResponse{
				Status: statusDown,
				Checks: []entity.Check{
					{
						Name:   "mysql",
						Status: statusDown,
						Error:  errors.New("error").Error(),
					},
					{
						Name:   "Redis",
						Status: statusUP,
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_mockDB, sqlMock, _ := mockDB.New("mysql")
			sqlMock.ExpectPing().WillReturnError(tt.dbErr)
			_mockCacheClient := &mockCacheClient.Client{}
			_mockCacheClient.On("Ping", mock.Anything).Return(tt.cacheErr)

			u := &useCase{
				startAt:     time.Now(),
				config:      cfg,
				db:          (*database.DB)(_mockDB),
				cacheClient: _mockCacheClient,
			}
			got, err := u.Readiness(tt.args.ctx)
			if !tt.wantErr(t, err, fmt.Sprintf("Readiness(%v)", tt.args.ctx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Readiness(%v)", tt.args.ctx)
		})
	}
}
