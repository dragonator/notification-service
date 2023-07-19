package notificationsending

import (
	"context"
)

const _notificationChannel = "notification_service_notifications"

// Operation provides an API for fetching single or multiple rentals.
type Operation struct {
	producer Producer
}

// NewOperation is a contruction function for Operation.
func NewOperation(publisher Producer) *Operation {
	return &Operation{
		producer: publisher,
	}
}

func (o *Operation) SendNotificationMessage(ctx context.Context, message string) error {
	return o.producer.Produce(ctx, _notificationChannel, message)
}
