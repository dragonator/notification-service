package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dragonator/notification-service/module/notification"
	"github.com/dragonator/notification-service/pkg/config"
	"github.com/dragonator/notification-service/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(cfg.LoggerLevel)

	consumerModule := notification.NewConsumerModule(cfg, logger)

	errs := make(chan error)
	consumerModule.NotificationConsumer.ConsumeMessages(context.Background(), errs)

	go func() {
		for err := range errs {
			logger.Errorf("error consuming message: %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop

	log.Printf("Signal caught (%s), stopping...", sig.String())
	consumerModule.NotificationConsumer.Stop()
	log.Print("Notification consumer stopped.")
}
