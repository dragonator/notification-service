package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dragonator/notification-service/module/notification/internal/event"
)

// ChannelHandler is a contract to a handler for a single message channel.
type ChannelHandler interface {
	SendMessage(ctx context.Context, message string) error
}

// Processor implements notification handling logic.
type Processor struct {
	channelHandlers []ChannelHandler
}

// NewProcessor is a constructior function for a Processor.
func NewProcessor(channelHandlers ...ChannelHandler) *Processor {
	return &Processor{
		channelHandlers: channelHandlers,
	}
}

// Process handles processing of a single notification message.
func (p *Processor) Process(ctx context.Context, msg []byte) error {
	notification := new(event.Notification)

	if err := json.Unmarshal(msg, notification); err != nil {
		return fmt.Errorf("unmarshling notification event: %w", err)
	}

	for _, c := range p.channelHandlers {
		if err := c.SendMessage(ctx, notification.Message); err != nil {
			return err
		}
	}

	return nil
}
