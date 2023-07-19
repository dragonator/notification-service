package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NotificationHandler is a contract to a notification handler.
type NotificationHandler interface {
	Notify(method, path string) func(w http.ResponseWriter, r *http.Request)
}

// NewRouter is a construction function for router that handles operations for notifications.
func NewRouter(nh NotificationHandler) http.Handler {
	router := chi.NewRouter()

	api := []struct {
		MethodFunc func(pattern string, handlerFn http.HandlerFunc)
		Method     string
		Path       string
		HandleFunc func(string, string) func(w http.ResponseWriter, r *http.Request)
	}{
		{router.Post, "POST", "/notify", nh.Notify},
	}

	for _, endpoint := range api {
		endpoint.MethodFunc(endpoint.Path, endpoint.HandleFunc(endpoint.Method, endpoint.Path))
	}

	return router
}
