package notification

import (
	"fmt"

	"github.com/dragonator/notification-service/module/notification/internal/http/handler"
	"github.com/dragonator/notification-service/module/notification/internal/http/service"
	"github.com/dragonator/notification-service/module/notification/internal/operation/notificationsending"
	"github.com/dragonator/notification-service/pkg/config"
	"github.com/dragonator/notification-service/pkg/logger"
)

// NotificationService provides methods for starting and stopping a notification service.
type NotificationService interface {
	Start()
	Stop()
}

// NotificationModule provides access to the functionality of notification module.
type NotificationModule struct {
	NotificationService NotificationService
}

// NewNotificationModule is a construction function for NotificationModule.
func NewNotificationModule(config *config.Config, logger *logger.Logger) (*NotificationModule, error) {
	notificationSendingOp := notificationsending.NewOperation(nil) // FIXME: Provide publisher
	notificationHandler := handler.NewNotificationHandler(notificationSendingOp)
	router := service.NewRouter(notificationHandler)

	notificationService, err := service.New(config, logger, router)
	if err != nil {
		return nil, fmt.Errorf("creating notification module: %w", err)
	}

	return &NotificationModule{
		NotificationService: notificationService,
	}, nil
}