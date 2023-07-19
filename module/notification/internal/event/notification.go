package event

// Notification is a kafka message contract used for notifications.
type Notification struct {
	Message string `json:"message,omitempty"`
}
