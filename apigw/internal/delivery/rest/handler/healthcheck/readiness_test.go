package healthcheck

import (
	"errors"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/health"
	"net/http"
	"net/http/httptest"
	"testing"

	mockUsecase "github.com/nofendian17/openota/apigw/internal/mocks/usecase/healthcheck"

	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Readiness(t *testing.T) {
	tests := []struct {
		name     string
		response []interface{}
		want     int
	}{
		{
			name: "should return 200",
			response: []interface{}{
				&health.ReadinessResponse{
					Status: "UP",
					Checks: []health.Check{
						{
							Name:   "mysql",
							Status: "UP",
							Error:  "",
						},
					},
				}, nil,
			},
			want: http.StatusOK,
		},
		{
			name: "should return 500",
			response: []interface{}{
				nil, errors.New("error"),
			},
			want: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mockUsecase.UseCase{}
			mockUseCase.On("Readiness", mock.Anything).Return(tt.response...)

			req := httptest.NewRequest(http.MethodGet, "/readiness", nil)
			req.Header.Set("Content-Type", "Application/json")
			rec := httptest.NewRecorder()
			defer rec.Flush()

			h := &handler{
				config:  config.New(),
				useCase: mockUseCase,
			}

			h.Readiness().ServeHTTP(rec, req)
			assert.Equal(t, tt.want, rec.Code)
		})
	}
}
