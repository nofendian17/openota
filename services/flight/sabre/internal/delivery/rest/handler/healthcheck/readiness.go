package healthcheck

import (
	"net/http"

	"github.com/nofendian17/openota/sabre/internal/delivery/rest/model/response"
)

func (h *handler) Readiness() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		res, err := h.useCase.Readiness(ctx)

		var status int
		var data interface{}
		var errs []error

		if err != nil {
			status = http.StatusInternalServerError
			errs = []error{err}
		} else {
			status = http.StatusOK
			data = res
		}

		httpResponse := response.New(
			status,
			http.StatusText(status),
			data,
			0,
			errs,
		)
		httpResponse.Json(w, status)
	}
}
