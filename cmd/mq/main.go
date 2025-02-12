package main

import (
	"fmt"
	"github.com/jekiapp/hi-mod-arch/internal/config"
	"github.com/jekiapp/hi-mod-arch/pkg/handler"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.InitConfig()

	application := initApplication(cfg)
	consumers := application.registerConsumer()

	stopChan := make(chan struct{})
	stopChannels := make([]chan struct{}, 0)

	for _, con := range consumers {
		stopChan, err := startConsumer(cfg.NsqConfig.NsqdAddress, con, stopChan)

		if err != nil {
			log.Fatalf("error start consumer: %s", err.Error())
			return
		}
		stopChannels = append(stopChannels, stopChan)
	}

	fmt.Println("All consumers started!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan // Wait for termination signal

	close(stopChan) // Signal consumers to stop
	for _, stopped := range stopChannels {
		if stopped != nil {
			<-stopped // Wait for each consumer to stop
		}
	}

	fmt.Println("All consumers finished.")
}

func startConsumer(nsqdAddress string, consumer handler.Consumer, stopChan chan struct{}) (chan struct{}, error) {
	nsqConsumer, err := nsq.NewConsumer(consumer.Topic, consumer.Channel, nsq.NewConfig())
	if err != nil {
		return nil, fmt.Errorf("consumer error:%s", err.Error())
	}

	nsqConsumer.AddHandler(consumer.Handler)
	if err := nsqConsumer.ConnectToNSQD(nsqdAddress); err != nil {
		return nil, fmt.Errorf("connection faield: %s", err.Error())
	}

	stopped := make(chan struct{})
	go func() {
		<-stopChan // Wait for stop signal
		nsqConsumer.Stop()
		<-nsqConsumer.StopChan // Wait for consumer to stop completely
		close(stopped)
	}()

	return stopped, nil
}
