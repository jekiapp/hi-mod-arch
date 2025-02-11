package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

// Define your message handler structure
type NsqHandler struct{}

// Implement the HandleMessage method for your handler
func (h *NsqHandler) HandleMessage(message *nsq.Message) error {
	// Log the message (or perform whatever logic you need)
	fmt.Printf("Received message: %s\n", string(message.Body))
	// Return nil if you successfully processed the message
	return nil
}

func initializeNsqConsumer(topic, channel, nsqAddr string) (*nsq.Consumer, error) {
	// Create a new NSQ consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create NSQ consumer: %w", err)
	}

	// Create an instance of your handler
	handler := &NsqHandler{}

	// Attach your handler to the consumer
	consumer.AddHandler(handler)

	// Connect the consumer to the NSQ daemon
	err = consumer.ConnectToNSQD(nsqAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NSQ daemon: %w", err)
	}

	// Return the consumer if everything is set up correctly
	return consumer, nil
}

func main() {
	// Initialize the consumer
	consumer, err := initializeNsqConsumer(topic, channel, nsqAddr)
	if err != nil {
		log.Fatalf("Error initializing NSQ consumer: %v", err)
	}

	// Wait until the consumer finishes its work
	select {
	case <-consumer.StopChan:
		fmt.Println("NSQ consumer stopped.")
	}
}
