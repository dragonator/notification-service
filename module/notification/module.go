package notification

import (
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/dragonator/notification-service/module/notification/internal/http/handler"
	"github.com/dragonator/notification-service/module/notification/internal/http/service"
	"github.com/dragonator/notification-service/module/notification/internal/operation/notificationsending"
	"github.com/dragonator/notification-service/module/notification/internal/processor/notification"
	"github.com/dragonator/notification-service/pkg/config"
	kafkaPKG "github.com/dragonator/notification-service/pkg/kafka"
	"github.com/dragonator/notification-service/pkg/logger"
)

const _notificationTopic = "notification_service_notifications"

// NotificationService provides methods for starting and stopping a notification service.
type NotificationService interface {
	Start()
	Stop()
}

// ServerModule provides access to the functionality of notification server module.
type ServerModule struct {
	NotificationService NotificationService
}

// NewServerModule is a construction function for ServerModule.
func NewServerModule(config *config.Config, logger *logger.Logger) (*ServerModule, error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.KafkaBrokerURL},
	})

	producer := kafkaPKG.NewProducer(w)
	notificationSendingOp := notificationsending.NewOperation(producer)
	notificationHandler := handler.NewNotificationHandler(notificationSendingOp)
	router := service.NewRouter(notificationHandler)

	notificationService, err := service.New(config, logger, router)
	if err != nil {
		return nil, fmt.Errorf("creating notification module: %w", err)
	}

	return &ServerModule{
		NotificationService: notificationService,
	}, nil
}

// ConsumerModule
type ConsumerModule struct {
	NotificationConsumer *kafkaPKG.Consumer
}

func NewConsumerModule(config *config.Config, logger *logger.Logger) *ConsumerModule {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.KafkaBrokerURL},
		GroupID:  "notification-service",
		Topic:    _notificationTopic,
		MaxBytes: 10e6, // 10MB
	})

	notificationProcessor := notification.NewProcessor(
		new(notification.SlackHandler),
		new(notification.SMSHandler),
		new(notification.EmailHandler),
	)

	return &ConsumerModule{
		NotificationConsumer: kafkaPKG.NewConsumer(r, notificationProcessor),
	}
}
