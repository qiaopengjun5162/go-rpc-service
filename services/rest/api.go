package rest

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/qiaopengjun5162/go-rpc-service/common/httputil"
	"github.com/qiaopengjun5162/go-rpc-service/config"
	"github.com/qiaopengjun5162/go-rpc-service/database"
	"github.com/qiaopengjun5162/go-rpc-service/services/rest/routes"
	"github.com/qiaopengjun5162/go-rpc-service/services/rest/service"
)

const (
	HealthPath          = "/health"
	SupportChainV1Path  = "/api/v1/support_chain"
	WalletAddressV1Path = "/api/v1/wallet_address"
)

type APIConfig struct {
	HTTPServer    config.ServerConfig
	MetricsServer config.ServerConfig
}

type API struct {
	router    *chi.Mux
	apiServer *httputil.HTTPServer
	db        *database.DB
	stopped   atomic.Bool
}

// NewApi initializes a new API instance from the given configuration.
//
// It creates a new API instance and calls `initFromConfig` to initialize the
// instance from the given configuration. If the initialization fails, it stops
// the API instance and joins the error with the stop error.
//
// Parameters:
//   - ctx: A context.Context that controls the initialization timeout.
//   - cfg: A pointer to the configuration to use for initialization.
//
// Returns:
//   - *API: A pointer to the initialized API instance if successful.
//   - error: An error if the initialization fails, or nil if successful.
func NewApi(ctx context.Context, cfg *config.Config) (*API, error) {
	out := &API{}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}

// initFromConfig initializes the API instance from the given configuration.
//
// It first calls `initDB` to initialize the database connection from the given
// configuration. If the initialization fails, it returns the error joined with
// the stop error.
//
// Then it calls `initRouter` to initialize the API router from the given
// configuration.
//
// Finally, it calls `startServer` to start the API server from the given
// configuration. If the start fails, it returns the error joined with the stop
// error.
//
// Parameters:
//   - ctx: A context.Context that controls the initialization timeout.
//   - cfg: A pointer to the configuration to use for initialization.
//
// Returns:
//   - error: An error if the initialization fails, or nil if successful.
func (a *API) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := a.initDB(ctx, cfg); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	a.initRouter(cfg.HTTPServer, cfg)
	if err := a.startServer(cfg.HTTPServer); err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	return nil
}

// initRouter initializes the API router with the specified server configuration.
//
// It creates a new instance of the Validator and HandleSrv to set up the service
// layer. A new chi router is created and configured with middleware for timeout,
// recovery, and heartbeat. It also sets up HTTP GET endpoints for support chain
// and wallet address using the provided routes.
//
// Parameters:
//   - conf: The server configuration for initializing the router.
//   - cfg: The application configuration used to set up the service.
func (a *API) initRouter(conf config.ServerConfig, cfg *config.Config) {
	v := new(service.Validator)

	svc := service.NewHandleSrv(v, a.db.Keys)
	apiRouter := chi.NewRouter()
	h := routes.NewRoutes(apiRouter, svc)

	apiRouter.Use(middleware.Timeout(time.Second * 12))
	apiRouter.Use(middleware.Recoverer)

	apiRouter.Use(middleware.Heartbeat(HealthPath))

	apiRouter.Get(fmt.Sprintf(SupportChainV1Path), h.GetSupportCoins)
	apiRouter.Get(fmt.Sprintf(WalletAddressV1Path), h.GetWalletAddress)

	a.router = apiRouter
}

// initDB initializes the database connection from the given configuration.
//
// It creates a new instance of the DB from the given configuration. If the
// initialization fails, it logs the error and returns it.
//
// Parameters:
//   - ctx: A context.Context that controls the initialization timeout.
//   - cfg: A pointer to the configuration to use for initialization.
//
// Returns:
//   - error: An error if the initialization fails, or nil if successful.
func (a *API) initDB(ctx context.Context, cfg *config.Config) error {
	initDb, err := database.NewDB(ctx, cfg.Database)
	if err != nil {
		log.Error("failed to connect to slave database", "err", err)
		return err
	}
	a.db = initDb
	return nil
}

// Start starts the API server.
//
// Parameters:
//   - ctx: A context.Context that controls the start timeout.
//
// Returns:
//   - error: An error if the start fails, or nil if successful.
func (a *API) Start(ctx context.Context) error {
	return nil
}

// Stop stops the API service.
//
// It stops the API server and closes the database connection.
// If any of the shutdown operations fail, it joins the errors together
// and returns the resulting error.
//
// Parameters:
//   - ctx: A context.Context that controls the shutdown timeout.
//
// Returns:
//   - error: An error if any of the shutdown operations fail, or nil if successful.
func (a *API) Stop(ctx context.Context) error {
	var result error
	if a.apiServer != nil {
		if err := a.apiServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop API server: %w", err))
		}
	}
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}
	a.stopped.Store(true)
	log.Info("API service shutdown complete")
	return result
}

// startServer starts the API server with the given server configuration.
//
// It constructs the server address from the provided host and port within
// the serverConfig. Then, it uses the router to initialize and start the
// HTTP server. If the server fails to start, it returns an error.
//
// Parameters:
//   - serverConfig: The ServerConfig containing the host and port for the API server.
//
// Returns:
//   - error: An error if the server fails to start, or nil if successful.
func (a *API) startServer(serverConfig config.ServerConfig) error {
	log.Debug("API server listening...", "port", serverConfig.Port)
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))
	srv, err := httputil.StarHttpServer(addr, a.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	log.Info("API server started", "addr", srv.Addr().String())
	a.apiServer = srv
	return nil
}

// Stopped returns a boolean indicating whether the API service has been stopped.
//
// It safely loads the stopped atomic boolean value and returns its value.
//
// Returns:
//   - bool: True if the API service has been stopped, false otherwise.
func (a *API) Stopped() bool {
	return a.stopped.Load()
}
