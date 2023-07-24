package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dragonator/notification-service/module/notification/internal/http/contract"
	"github.com/dragonator/notification-service/module/notification/internal/http/service/svc"
)

// NotificationSendingOp is a contract to a notification sending operation.
type NotificationSendingOp interface {
	SendNotificationMessage(ctx context.Context, message string) error
}

// NotificationHandler holds implementation of handlers for rentals.
type NotificationHandler struct {
	notificationSendingOp NotificationSendingOp
}

// NewNotificationHandler is a construction function for NotificationHandler.
func NewNotificationHandler(notificationSendingOp NotificationSendingOp) *NotificationHandler {
	return &NotificationHandler{
		notificationSendingOp: notificationSendingOp,
	}
}

// Notify sends a notification message to be propagated to appropriate channels.
func (rh *NotificationHandler) Notify(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &contract.NotificationRequest{}
		if err := decode(r, req); err != nil {
			errorResponse(w, fmt.Errorf("%w: %v", svc.ErrDecodeRequest, err))
			return
		}

		err := rh.notificationSendingOp.SendNotificationMessage(r.Context(), req.Message)
		if err != nil {
			errorResponse(w, err)
			return
		}

		successResponse(w, nil)

		return
	}
}
