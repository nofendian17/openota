package healthcheck

import (
	"net/http"

	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
)

func (h *handler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		res, err := h.useCase.Health(ctx)

		var status int
		var message string
		var data interface{}

		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		} else {
			status = http.StatusOK
			message = http.StatusText(status)
			data = res
		}

		httpResponse := response.New(status, message, data, 0, nil)
		httpResponse.Json(w, status)
	}
}
