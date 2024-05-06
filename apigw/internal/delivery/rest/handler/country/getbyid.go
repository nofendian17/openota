package country

import (
	"errors"
	"net/http"

	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/country"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"gorm.io/gorm"
)

func (h *handler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID := r.PathValue("countryID")

		// Validate request
		if err := h.validator.Validate(&country.GetByID{ID: ID}); err != nil {
			httpResponse := response.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, 0, err)
			httpResponse.Json(w, http.StatusBadRequest)
			return
		}

		// Get country by ID
		res, err := h.useCase.GetByID(ctx, ID)

		// Handle response
		var status int
		var message string
		var data interface{}

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
	}
}
