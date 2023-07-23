package notification

import (
	"context"
	"fmt"
)

type SlackHandler struct{}

func (sh *SlackHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via Slack: %s\n", message)
	return nil
}

type EmailHandler struct{}

func (eh *EmailHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via Email: %s\n", message)
	return nil
}

type SMSHandler struct{}

func (smsh *SMSHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via SMS: %s\n", message)
	return nil
}
