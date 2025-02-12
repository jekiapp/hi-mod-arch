package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type GenericHandlerHttp interface {
	// the ideal signature would be having context as the first parameter
	// I omitted it to maintain simplicity
	HandlerFunc(input interface{}) (output interface{}, err error)
	ObjectAddress() interface{}
}

func HttpGenericHandler(handler GenericHandlerHttp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		data := handler.ObjectAddress()
		if err := json.Unmarshal(body, data); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// (optional) validate input data using json validator

		result, err := handler.HandlerFunc(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var resp responseHttpTemplate
		if err != nil {
			resp = responseHttpTemplate{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp = responseHttpTemplate{
			Status:  "ok",
			Message: "succes",
			Data:    result,
		}

		json.NewEncoder(w).Encode(resp)
	}
}

type responseHttpTemplate struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
