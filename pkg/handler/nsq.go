package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type NsqHandlerResult struct {
	Requeue time.Duration
	Finish  bool
}

type GenericHandlerNsq interface {
	// the ideal signature would be having context as the first parameter
	// I omitted it to maintain simplicity
	HandlerFunc(input interface{}) (output NsqHandlerResult, err error)
	ObjectAddress() interface{}
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

type Consumer struct {
	Topic   string
	Channel string
	Config  *nsq.Config
	Handler nsq.Handler
}

func NewGenericConsumer(topic, channel string, config *nsq.Config, handler GenericHandlerNsq) Consumer {
	return Consumer{
		Topic:   topic,
		Channel: channel,
		Config:  config,
		Handler: NsqGenericHandler(handler),
	}
}
