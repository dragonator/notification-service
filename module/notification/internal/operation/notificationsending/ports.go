package notificationsending

import (
	"context"
)

// Publisher is a contract to a message publisher.
type Publisher interface {
	Publish(ctx context.Context, channel string, message any) error
}
