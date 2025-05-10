package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type GenericHandlerHttp[I any, O any] interface {
	HandlerFunc(ctx context.Context, input I) (output O, err error)
}

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
)

type Response[O any] struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    O              `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
}

func HttpGenericHandler[I any, O any](handler GenericHandlerHttp[I, O]) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set content type header
		w.Header().Set("Content-Type", "application/json")

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[O]{
				Status:  StatusError,
				Message: "Failed to read request body",
				Error:   err.Error(),
			})
			return
		}
		defer r.Body.Close()

		// Parse request body
		data := new(I)
		if err := json.Unmarshal(body, data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[O]{
				Status:  StatusError,
				Message: "Invalid JSON format",
				Error:   err.Error(),
			})
			return
		}

		// Validate input data
		if err := validate.Struct(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[O]{
				Status:  StatusError,
				Message: "Validation failed",
				Error:   err.Error(),
			})
			return
		}

		// Execute handler
		result, err := handler.HandlerFunc(r.Context(), *data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response[O]{
				Status:  StatusError,
				Message: "Handler execution failed",
				Error:   err.Error(),
			})
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response[O]{
			Status:  StatusSuccess,
			Message: "Operation completed successfully",
			Data:    result,
		})
	}
}
