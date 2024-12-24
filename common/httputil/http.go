package httputil

import (
	"net/http"
)

var timeouts = DefaultTimeOuts

// NewHttpServer creates a new HTTP server from the given handler,
// with default timeouts for reads, writes, and idle connections.
func NewHttpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
	}
}
