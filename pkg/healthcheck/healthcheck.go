// Package healthcheck provides a simple way to define health checks for external services or components.
// These checks are aggregated by the /status endpoint.
package healthcheck

import (
	"context"
)

// HealthChecker is the interface that wraps the HealthCheck method.
type HealthChecker interface {
	HealthCheck(ctx context.Context) error
}

// HealthCheck is a structure containing the configuration for a single health check.
type HealthCheck struct {
	// ID is a unique identifier for the healthcheck.
	ID string

	// Checker is the function used to perform the healthchecks.
	Checker HealthChecker
}

// New creates a new instance of a health check configuration with default timeout.
func New(id string, checker HealthChecker) HealthCheck {
	return HealthCheck{
		ID:      id,
		Checker: checker,
	}
}
