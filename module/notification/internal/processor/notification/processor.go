package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dragonator/notification-service/module/notification/internal/event"
)

type ChannelHandler interface {
	SendMessage(ctx context.Context, message string) error
}

type Processor struct {
	channelHandlers []ChannelHandler
}

func NewProcessor(channelHandlers ...ChannelHandler) *Processor {
	return &Processor{
		channelHandlers: channelHandlers,
	}
}

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
