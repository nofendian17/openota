package city

import (
	"encoding/json"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response"
	"net/http"
)

func (h *handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Decode request
		decoder := json.NewDecoder(r.Body)
		var requestBody city.Create
		if err := decoder.Decode(&requestBody); err != nil {
			httpResponse := response.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, 0, err)
			httpResponse.Json(w, http.StatusBadRequest)
			return
		}

		// Validate request
		if err := h.validator.Validate(&requestBody); err != nil {
			httpResponse := response.New(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil, 0, err)
			httpResponse.Json(w, http.StatusUnprocessableEntity)
			return
		}

		err := h.useCase.Create(ctx, requestBody)

		var (
			// Handle response
			status  int
			message string
		)

		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		} else {
			status = http.StatusCreated
			message = http.StatusText(status)
		}

		httpResponse := response.New(status, message, nil, 0, nil)
		httpResponse.Json(w, status)
		return
	}
}
