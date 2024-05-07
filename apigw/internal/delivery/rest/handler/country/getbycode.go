package country

import (
	"errors"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"gorm.io/gorm"
	"net/http"
)

func (h *handler) GetByCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		countryCode := r.PathValue("countryCode")
		res, err := h.useCase.GetByCode(ctx, countryCode)

		var (
			status  int
			message string
			data    interface{}
		)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				status = http.StatusNotFound
				message = err.Error()
			} else {
				status = http.StatusInternalServerError
				message = err.Error()
			}
		} else {
			status = http.StatusOK
			message = http.StatusText(status)
			data = res
		}

		httpResponse := response.New(status, message, data, 0, nil)
		httpResponse.Json(w, status)
		return
	}
}
