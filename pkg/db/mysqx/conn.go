// Package mysqx provides database access
package mysqx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	defaultConnArgs = "parseTime=true&loc=Local"
)

type DBGetter interface {
	Get(database string) *sql.DB
}

type Options struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func (c *Options) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true&charset=utf8mb4&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}

type Config struct {
	Driver        string
	ConnectionURL string

	MaxIdleConns int
	MaxOpenConns int
}

type Connection struct {
	config *Config
	conn   *sql.DB
}

func humanError(err error, cfg Config) error {
	return fmt.Errorf("failed to open mysql, driver: %s, error: %w",
		cfg.Driver, err)
}

func NewConnectionWithOptions(options Options) (*Connection, error) {
	return NewConnection(Config{
		Driver:        "mysql",
		ConnectionURL: options.DSN(),
	})
}

func NewConnection(cfg Config) (*Connection, error) {
	config := cfg
	if err := mergo.Merge(&config, defaultConfiguration()); err != nil {
		return nil, errors.Wrap(err, "failed to merge configuration")
	}

	conn, err := sqlx.Open(cfg.Driver, config.ConnectionURL)
	if err != nil {
		return nil, humanError(err, config)
	}

	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetMaxIdleConns(config.MaxIdleConns)

	return &Connection{
		config: &config,
		conn:   conn.DB,
	}, nil
}

func (c *Connection) Ping(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := c.conn.PingContext(ctx); err != nil {
		return humanError(err, *c.config)
	}

	return nil
}

func defaultConfiguration() Config {
	return Config{
		Driver:        "mysql",
		ConnectionURL: "root@(localhost)/dev?" + defaultConnArgs,
		MaxIdleConns:  0,
		MaxOpenConns:  20,
	}
}

func (c *Connection) GetDB() *sql.DB {
	return c.conn
}
