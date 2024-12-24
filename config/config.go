package config

import (
	"github.com/qiaopengjun5162/go-rpc-service/flags"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Migrations    string
	Database      DBConfig
	RpcServer     ServerConfig
	HTTPServer    ServerConfig
	MetricsServer ServerConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

// NewConfig creates a new instance of Config from the given CLI context.
//
// It extracts settings for database, RPC server, HTTP server, and metrics server
// from the provided CLI context flags. These settings include host, port, name,
// user, and password for the database, as well as host and port for the servers.
//
// Parameters:
//   - ctx: A cli.Context containing the CLI flag values.
//
// Returns:
//   - Config: A Config instance populated with values extracted from the CLI context.
func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
		Database: DBConfig{
			Host:     ctx.String(flags.DbHostFlag.Name),
			Port:     ctx.Int(flags.DbPortFlag.Name),
			Name:     ctx.String(flags.DbNameFlag.Name),
			User:     ctx.String(flags.DbUserFlag.Name),
			Password: ctx.String(flags.DbPasswordFlag.Name),
		},
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
		HTTPServer: ServerConfig{
			Host: ctx.String(flags.HttpHostFlag.Name),
			Port: ctx.Int(flags.HttpPortFlag.Name),
		},
		MetricsServer: ServerConfig{
			Host: ctx.String(flags.MetricsHostFlag.Name),
			Port: ctx.Int(flags.MetricsPortFlag.Name),
		},
	}
}
