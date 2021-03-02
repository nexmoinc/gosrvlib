package sqlconn

import (
	"time"
)

// Option is a type alias for a function that configures the DB connector.
type Option func(*config)

// WithQuoteIDFunc replaces the default QuoteID function.
func WithQuoteIDFunc(fn SQLQuoteFunc) Option {
	return func(cfg *config) {
		cfg.quoteIDFunc = fn
	}
}

// WithQuoteValueFunc replaces the default QuoteValue function.
func WithQuoteValueFunc(fn SQLQuoteFunc) Option {
	return func(cfg *config) {
		cfg.quoteValueFunc = fn
	}
}

// WithConnectFunc replaces the default connection function.
func WithConnectFunc(fn ConnectFunc) Option {
	return func(cfg *config) {
		cfg.connectFunc = fn
	}
}

// WithCheckConnectionFunc replaces the default connection check function.
func WithCheckConnectionFunc(fn CheckConnectionFunc) Option {
	return func(cfg *config) {
		cfg.checkConnectionFunc = fn
	}
}

// WithSQLOpenFunc replaces the default open database function.
func WithSQLOpenFunc(fn SQLOpenFunc) Option {
	return func(cfg *config) {
		cfg.sqlOpenFunc = fn
	}
}

// WithConnectMaxRetry sets the maximum retry attempts.
func WithConnectMaxRetry(maxRetry int) Option {
	return func(cfg *config) {
		cfg.connectMaxRetry = maxRetry
	}
}

// WithConnectRetryInterval sets the interval between connection retries.
func WithConnectRetryInterval(interval time.Duration) Option {
	return func(cfg *config) {
		cfg.connectRetryInterval = interval
	}
}

// WithConnMaxLifetime sets the maximum lifetime of a database connection.
func WithConnMaxLifetime(lifetime time.Duration) Option {
	return func(cfg *config) {
		cfg.connMaxLifetime = lifetime
	}
}

// WithConnMaxIdle sets the maximum number of idle database connections.
func WithConnMaxIdle(maxIdle int) Option {
	return func(cfg *config) {
		cfg.connMaxIdle = maxIdle
	}
}

// WithConnMaxOpen sets the maximum number of open database connections.
func WithConnMaxOpen(maxOpen int) Option {
	return func(cfg *config) {
		cfg.connMaxOpen = maxOpen
	}
}

// WithDefaultDriver sets the default driver to use if not included in the DSN.
func WithDefaultDriver(driver string) Option {
	return func(cfg *config) {
		if cfg.driver == "" {
			cfg.driver = driver
		}
	}
}
