package notificationsending

import (
	"context"
)

// Producer is a contract to a message producer.
type Producer interface {
	Produce(ctx context.Context, channel string, message any) error
}
