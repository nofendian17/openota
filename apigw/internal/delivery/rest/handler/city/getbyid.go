package city

import (
	"errors"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"gorm.io/gorm"
	"net/http"
)

func (h *handler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID := r.PathValue("cityID")

		// Validate request
		if err := h.validator.Validate(&city.GetByID{ID: ID}); err != nil {
			httpResponse := response.New(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil, 0, err)
			httpResponse.Json(w, http.StatusUnprocessableEntity)
			return
		}

		// Get state by ID
		res, err := h.useCase.GetByID(ctx, ID)

		var (
			// Handle response
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
