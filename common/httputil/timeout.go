package httputil

import "time"

var DefaultTimeOuts = HTTPTimeouts{
	ReadTimeout:       time.Second * 15,
	ReadHeaderTimeout: time.Second * 15,
	WriteTimeout:      time.Second * 15,
	IdleTimeout:       time.Second * 120,
}

type HTTPTimeouts struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

// WithTimeouts returns an HTTPOption that sets the read, read header, write,
// and idle timeouts on an HTTPServer using the provided HTTPTimeouts struct.
func WithTimeouts(timeouts HTTPTimeouts) HTTPOption {
	return func(s *HTTPServer) error {
		s.srv.ReadTimeout = timeouts.ReadTimeout
		s.srv.ReadHeaderTimeout = timeouts.ReadHeaderTimeout
		s.srv.WriteTimeout = timeouts.WriteTimeout
		s.srv.IdleTimeout = timeouts.IdleTimeout
		return nil
	}
}
