package city

import (
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/state"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"net/http"
)

func (h *handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID := r.PathValue("cityID")

		var requestBody state.Delete

		// Validate request
		requestBody.ID = ID
		if err := h.validator.Validate(&requestBody); err != nil {
			httpResponse := response.New(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil, 0, err)
			httpResponse.Json(w, http.StatusUnprocessableEntity)
			return
		}

		err := h.useCase.Delete(ctx, ID)

		var (
			// Handle response
			status  int
			message string
		)

		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		} else {
			status = http.StatusOK
			message = http.StatusText(status)
		}

		httpResponse := response.New(status, message, nil, 0, nil)
		httpResponse.Json(w, status)
		return
	}
}
