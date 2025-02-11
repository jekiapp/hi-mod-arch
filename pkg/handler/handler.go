package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"io"
	"net/http"
	"time"
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

type GenericHandlerNsq interface {
	// the ideal signature would be having context as the first parameter
	// I omitted it to maintain simplicity
	HandlerFunc(input interface{}) (output NsqHandlerResult, err error)
	ObjectAddress() interface{}
}
type responseHttpTemplate struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NsqHandlerResult struct {
	Requeue time.Duration
	Finish  bool
}

func NsqGenericHandler(handler GenericHandlerNsq) nsq.HandlerFunc {
	return func(msg *nsq.Message) error {
		body := msg.Body
		data := handler.ObjectAddress()
		if err := json.Unmarshal(body, data); err != nil {
			return fmt.Errorf("error unmarshal object %+v", data)
		}

		// (optional) validate input object using json validator

		output, err := handler.HandlerFunc(data)
		if err != nil {
			if output.Requeue != 0 {
				msg.Requeue(output.Requeue)
			} else if output.Finish {
				msg.Finish()
			}

			return err
		}

		msg.Finish()
		return nil
	}

}
