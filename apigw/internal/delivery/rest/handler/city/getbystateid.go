package city

import (
	"errors"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"gorm.io/gorm"
	"net/http"
)

func (h *handler) GetByStateID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID := r.PathValue("stateID")

		// Validate request
		if err := h.validator.Validate(&city.GetByStateID{ID: ID}); err != nil {
			httpResponse := response.New(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil, 0, err)
			httpResponse.Json(w, http.StatusUnprocessableEntity)
			return
		}
		res, err := h.useCase.GetByStateID(ctx, ID)

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
