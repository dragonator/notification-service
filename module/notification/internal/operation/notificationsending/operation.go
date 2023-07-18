package notificationsending

import (
	"context"
)

const _notificationChannel = "notification_service_notifications"

// Operation provides an API for fetching single or multiple rentals.
type Operation struct {
	publisher Publisher
}

// NewOperation is a contruction function for Operation.
func NewOperation(publisher Publisher) *Operation {
	return &Operation{
		publisher: publisher,
	}
}

func (o *Operation) SendNotificationMessage(ctx context.Context, message string) error {
	return o.publisher.Publish(ctx, _notificationChannel, message)
}
