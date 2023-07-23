package notification

import (
	"context"
	"fmt"
)

// SlackHandler holds logic for sending a message via Slack.
type SlackHandler struct{}

// SendMessage simulates sending a message via Slack.
func (sh *SlackHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via Slack: %s\n", message)
	return nil
}

// SlackHandler holds logic for sending a message via Email.
type EmailHandler struct{}

// SendMessage simulates sending a message via Email.
func (eh *EmailHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via Email: %s\n", message)
	return nil
}

// SlackHandler holds logic for sending a message via SMS.
type SMSHandler struct{}

// SendMessage simulates sending a message via SMS.
func (smsh *SMSHandler) SendMessage(ctx context.Context, message string) error {
	fmt.Printf("Sending message via SMS: %s\n", message)
	return nil
}
