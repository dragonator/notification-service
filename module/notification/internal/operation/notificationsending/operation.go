package notificationsending

import (
	"context"

	"github.com/dragonator/notification-service/module/notification/internal/event"
)

const _notificationChannel = "notification_service_notifications"

// Operation provides an API for fetching single or multiple notifications.
type Operation struct {
	producer Producer
}

// NewOperation is a contruction function for Operation.
func NewOperation(publisher Producer) *Operation {
	return &Operation{
		producer: publisher,
	}
}

// SendNotificationMessage sends the given message as kafka event.
func (o *Operation) SendNotificationMessage(ctx context.Context, message string) error {
	return o.producer.Produce(ctx, _notificationChannel, &event.Notification{Message: message})
}
