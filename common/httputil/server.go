package httputil

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"github.com/pkg/errors"
)

type HTTPServer struct {
	listener net.Listener
	srv      *http.Server
	closed   atomic.Bool
}

type HTTPOption func(svr *HTTPServer) error

// StarHttpServer starts a new HTTP server with the given address, handler, and options.
//
// It creates a TCP listener, initializes an http.Server with the provided handler and default timeouts,
// and applies any additional options. The server is started in a separate goroutine.
//
// Parameters:
//   - addr: A string representing the network address to listen on (e.g., ":8080").
//   - handler: An http.Handler to handle incoming HTTP requests.
//   - opts: Optional variadic HTTPOption functions to customize the server configuration.
//
// Returns:
//   - *HTTPServer: A pointer to the created HTTPServer instance if successful.
//   - error: An error if the server initialization fails, or nil if successful.
func StarHttpServer(addr string, handler http.Handler, opts ...HTTPOption) (*HTTPServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen error=", err)
		return nil, errors.New("Init listener fail")
	}
	srvCtx, srvCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			return srvCtx
		},
	}
	out := &HTTPServer{listener: listener, srv: srv}

	for _, opt := range opts {
		if err := opt(out); err != nil {
			srvCancel()
			fmt.Println("apply err:", err)
			return nil, errors.New("One of http op fail")
		}
	}
	go func() {
		err := out.srv.Serve(listener)
		srvCancel()
		if errors.Is(err, http.ErrServerClosed) {
			out.closed.Store(true)
		} else {
			fmt.Println("unknown err:", err)
			panic("unknown error")
		}
	}()
	return out, nil
}

// Closed returns true if the server has been stopped.
//
// This method is safe to call concurrently.
func (hs *HTTPServer) Closed() bool {
	return hs.closed.Load()
}

// Stop gracefully stops the HTTP server.
//
// It first attempts to shut down the server using the provided context.
// If the shutdown fails due to a context deadline or cancellation,
// it forcefully closes the server.
//
// Parameters:
//   - ctx: A context.Context that controls the shutdown timeout.
//
// Returns:
//   - error: An error if the shutdown or close operation fails, or nil if successful.
func (hs *HTTPServer) Stop(ctx context.Context) error {
	if err := hs.Shutdown(ctx); err != nil {
		if errors.Is(err, ctx.Err()) {
			return hs.Close()
		}
		return err
	}
	return nil
}

func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	return hs.srv.Shutdown(ctx)
}

func (hs *HTTPServer) Close() error {
	return hs.srv.Close()
}

func (hs *HTTPServer) Addr() net.Addr {
	return hs.listener.Addr()
}

// WithMaxHeaderBytes returns an HTTPOption that sets the maximum allowed size for request headers (including the request line).
//
// Parameters:
//   - max: The maximum allowed size in bytes.
//
// Returns:
//   - HTTPOption: An HTTPOption that can be used to configure an HTTPServer.
func WithMaxHeaderBytes(max int) HTTPOption {
	return func(srv *HTTPServer) error {
		srv.srv.MaxHeaderBytes = max
		return nil
	}
}
