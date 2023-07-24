package main

import (
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

	serverModule, err := notification.NewServerModule(cfg, logger)
	if err != nil {
		panic(err)
	}

	serverModule.NotificationService.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop

	log.Printf("Signal caught (%s), stopping...", sig.String())
	serverModule.NotificationService.Stop()
	log.Print("Service stopped.")
}
