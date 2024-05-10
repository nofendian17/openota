package state

import (
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"net/http"
)

func (h *handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		res, err := h.useCase.GetAll(ctx)

		var (
			status  int
			data    interface{}
			message string
		)

		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		} else {
			status = http.StatusOK
			message = http.StatusText(status)
			data = res
		}

		httpResponse := response.New(status, message, data, int64(len(res)), nil)
		httpResponse.Json(w, status)
		return
	}
}
