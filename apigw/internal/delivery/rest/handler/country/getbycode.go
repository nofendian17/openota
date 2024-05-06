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

		var status int
		var data interface{}
		var errs []error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				status = http.StatusNotFound
			} else {
				status = http.StatusInternalServerError
			}
			errs = []error{err}
		} else {
			status = http.StatusOK
			data = res
		}

		httpResponse := response.New(status, http.StatusText(status), data, 0, errs)
		httpResponse.Json(w, status)
	}
}
