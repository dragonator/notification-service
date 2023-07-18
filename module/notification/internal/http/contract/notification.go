package contract

// NotificationRequest is a client request object for sending notifications.
type NotificationRequest struct {
	Message string `json:"message"`
}
